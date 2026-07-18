package services

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Payment struct {
	DB          *sql.DB
	ChapaSecret string
}


var chapaHTTPClient = &http.Client{
	Timeout: 15 * time.Second,
}


type chapaInitRequest struct {
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	TxRef       string `json:"tx_ref"`
	CallbackURL string `json:"callback_url"`
	ReturnURL   string `json:"return_url"`
}

type chapaInitResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`

	Data struct {
		CheckoutURL string `json:"checkout_url"`
	} `json:"data"`
}

type chapaVerifyResponse struct {
	Status string `json:"status"`
	Data   struct {
		Status string `json:"status"`
	} `json:"data"`
}


type createBookingPayload struct {
	Input struct {
		ScheduleID uuid.UUID   `json:"scheduleId"`
		SeatIDs    []uuid.UUID `json:"seatIds"`
	} `json:"input"`

	SessionVariables map[string]string `json:"session_variables"`
}

func (p *Payment) CreateBookingAction(c *gin.Context) {

	var payload struct {
		Input struct {
			ScheduleID      uuid.UUID   `json:"schedule_id"`
			ScheduleSeatIDs []uuid.UUID `json:"schedule_seat_ids"`
		} `json:"input"`

		SessionVariables map[string]string `json:"session_variables"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID := payload.SessionVariables["x-hasura-user-id"]

	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "missing user session",
		})
		return
	}

	if len(payload.Input.ScheduleSeatIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no seats selected",
		})
		return
	}

	tx, err := p.DB.Begin()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to start transaction",
		})
		return
	}

	defer tx.Rollback()

	bookingRef := "book-" + uuid.New().String()

	var ticketIDs []uuid.UUID

	for _, seatID := range payload.Input.ScheduleSeatIDs {

		var isAvailable bool

		err := tx.QueryRow(
			`
			SELECT is_available
			FROM schedule_seats
			WHERE id=$1
			FOR UPDATE
			`,
			seatID,
		).Scan(&isAvailable)

		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "seat not found",
				})
				return
			}

			log.Println("seat lock error:", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to check seat availability",
			})
			return
		}

		if !isAvailable {
			c.JSON(http.StatusConflict, gin.H{
				"message": "one or more seats are no longer available",
			})
			return
		}

		_, err = tx.Exec(
			`
			UPDATE schedule_seats
			SET is_available=false
			WHERE id=$1
			`,
			seatID,
		)

		if err != nil {
			log.Println("seat update error:", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to lock seat",
			})
			return
		}

		var ticketID uuid.UUID

		err = tx.QueryRow(
			`
			INSERT INTO tickets
			(
				user_id,
				schedule_id,
				schedule_seat_id,
				status,
				booking_reference
			)
			VALUES
			(
				$1,
				$2,
				$3,
				'pending',
				$4
			)
			RETURNING id
			`,
			userID,
			payload.Input.ScheduleID,
			seatID,
			bookingRef,
		).Scan(&ticketID)

		if err != nil {
			log.Println("ticket insert error:", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to create ticket",
			})
			return
		}

		ticketIDs = append(ticketIDs, ticketID)
	}

	if err := tx.Commit(); err != nil {
		log.Println("booking commit error:", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to commit booking",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":           true,
		"message":           "Booking created successfully",
		"booking_reference": bookingRef,
		"ticket_ids":        ticketIDs,
		"payment_id":        nil,
		"amount":            nil,
	})
}


type initiatePaymentPayload struct {

	Input struct {
		BookingReference string `json:"booking_reference"`
	} `json:"input"`

	SessionVariables map[string]string `json:"session_variables"`
}


func (p *Payment) InitiatePaymentAction(c *gin.Context) {

	var payload initiatePaymentPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID := payload.SessionVariables["x-hasura-user-id"]

	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "missing user session",
		})
		return
	}

	bookingRef := payload.Input.BookingReference

	if bookingRef == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "missing booking reference",
		})
		return
	}


	var email string
	var firstName string

	err := p.DB.QueryRow(
		`
		SELECT email, name
		FROM users
		WHERE id=$1
		`,
		userID,
	).Scan(
		&email,
		&firstName,
	)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {

			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "user not found",
			})
			return
		}

		log.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get user information",
		})

		return
	}


	var amount string
	var ticketCount int

	err = p.DB.QueryRow(
		`
		SELECT COALESCE(SUM(s.price), 0)::integer::text, COUNT(t.id)
		FROM tickets t
		JOIN schedules s
		ON s.id=t.schedule_id
		WHERE t.booking_reference=$1
		AND t.user_id=$2
		AND t.status='pending'
		`,
		bookingRef,
		userID,
	).Scan(&amount, &ticketCount)

	if err != nil {
		log.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get booking information",
		})

		return
	}

	if ticketCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "booking not found or already paid",
		})
		return
	}


	txRef := "cinema-" + uuid.New().String()


	var paymentID uuid.UUID

	err = p.DB.QueryRow(
		`
		INSERT INTO payments
		(
			user_id,
			booking_reference,
			transaction_ref,
			status,
			amount,
			currency
		)
		VALUES
		(
			$1,
			$2,
			$3,
			'pending',
			$4,
			'ETB'
		)
		RETURNING id
		`,
		userID,
		bookingRef,
		txRef,
		amount,
	).Scan(&paymentID)

	if err != nil {
		log.Println("payment insert error:", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}


	chapaRequest := chapaInitRequest{

		Amount: amount,

		Currency: "ETB",

		Email: email,

		FirstName: firstName,

		LastName: "",

		TxRef: txRef,

		CallbackURL: "https://carpentry-wisplike-agonizing.ngrok-free.dev/payment/webhook",


		ReturnURL: fmt.Sprintf(
			"http://localhost:3000/payment/success?trx_ref=%s&booking_reference=%s",
			txRef,
			bookingRef,
		),
	}

	body, err := json.Marshal(chapaRequest)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to create chapa request",
		})

		return
	}

	request, err := http.NewRequest(
		"POST",
		"https://api.chapa.co/v1/transaction/initialize",
		bytes.NewBuffer(body),
	)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to create request",
		})

		return
	}

	request.Header.Set(
		"Authorization",
		"Bearer "+p.ChapaSecret,
	)

	request.Header.Set(
		"Content-Type",
		"application/json",
	)

	response, err := chapaHTTPClient.Do(request)

	if err != nil {
		log.Println("chapa initialize request error:", err)

		c.JSON(502, gin.H{
			"message": "failed to connect to Chapa",
		})
		return
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed reading Chapa response",
		})
		return
	}

	log.Println("Chapa HTTP status:", response.Status)
	log.Println("Chapa response:", string(responseBody))

	var chapaResponse chapaInitResponse

	err = json.Unmarshal(responseBody, &chapaResponse)

	if err != nil {
		c.JSON(500, gin.H{
			"message":  "invalid chapa response",
			"response": string(responseBody),
		})

		return
	}

	if chapaResponse.Status != "success" {
		c.JSON(502, gin.H{
			"message": fmt.Sprintf(
				"Chapa error: %s",
				chapaResponse.Message,
			),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{

		"checkoutUrl": chapaResponse.Data.CheckoutURL,

		"txRef": txRef,
	})
}


func extractTxRef(c *gin.Context, jsonTxRef string) string {
	if jsonTxRef != "" {
		return jsonTxRef
	}

	if v := c.PostForm("tx_ref"); v != "" {
		return v
	}

	if v := c.Query("tx_ref"); v != "" {
		return v
	}

	if v := c.Query("trx_ref"); v != "" {
		return v
	}

	return ""
}

func (p *Payment) PaymentWebhook(c *gin.Context) {

	var payload struct {
		TxRef string `json:"tx_ref"`

		Status string `json:"status"`
	}


	_ = c.ShouldBindJSON(&payload)

	txRef := extractTxRef(c, payload.TxRef)

	if txRef == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing tx_ref",
		})
		return
	}

	newStatus, err := p.verifyWithChapa(txRef)

	if err != nil {
		log.Println("webhook verify error:", err)

		c.JSON(502, gin.H{
			"error": "verification failed",
		})

		return
	}

	if err := p.applyPaymentStatus(txRef, newStatus); err != nil {
		log.Println("webhook apply status error:", err)

		c.JSON(500, gin.H{
			"error": "failed to update payment",
		})

		return
	}

	c.JSON(200, gin.H{
		"received": true,
	})
}

func (p *Payment) verifyWithChapa(txRef string) (string, error) {

	verifyRequest, err := http.NewRequest(
		"GET",
		"https://api.chapa.co/v1/transaction/verify/"+txRef,
		nil,
	)

	if err != nil {
		return "", fmt.Errorf("building verify request: %w", err)
	}

	verifyRequest.Header.Set(
		"Authorization",
		"Bearer "+p.ChapaSecret,
	)

	response, err := chapaHTTPClient.Do(verifyRequest)

	if err != nil {
		return "", fmt.Errorf("calling chapa verify: %w", err)
	}

	defer response.Body.Close()

	verifyBody, err := io.ReadAll(response.Body)

	if err != nil {
		return "", fmt.Errorf("reading verify response: %w", err)
	}

	log.Println("Chapa verify HTTP status:", response.Status)
	log.Println("Chapa verify response:", string(verifyBody))

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return "", fmt.Errorf(
			"chapa verify returned status %d",
			response.StatusCode,
		)
	}

	var verifyResponse chapaVerifyResponse

	if err := json.Unmarshal(verifyBody, &verifyResponse); err != nil {
		return "", fmt.Errorf("decoding verify response: %w", err)
	}

	if verifyResponse.Status == "success" && verifyResponse.Data.Status == "success" {
		return "success", nil
	}

	return "failed", nil
}


func (p *Payment) applyPaymentStatus(txRef string, newStatus string) error {

	tx, err := p.DB.Begin()

	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	defer tx.Rollback()

	_, err = tx.Exec(
		`
		UPDATE payments
		SET status=$1,
		updated_at=NOW()
		WHERE transaction_ref=$2
		`,
		newStatus,
		txRef,
	)

	if err != nil {
		return fmt.Errorf("updating payment: %w", err)
	}


	ticketStatus := "confirmed"

	if newStatus != "success" {
		ticketStatus = "failed"
	}

	_, err = tx.Exec(
		`
		UPDATE tickets
		SET status=$1,
		updated_at=NOW()
		WHERE booking_reference = (
			SELECT booking_reference
			FROM payments
			WHERE transaction_ref=$2
		)
		AND status='pending'
		`,
		ticketStatus,
		txRef,
	)

	if err != nil {
		return fmt.Errorf("updating tickets: %w", err)
	}


	if newStatus != "success" {

		_, err = tx.Exec(
			`
			UPDATE schedule_seats
			SET is_available=true
			WHERE id IN (
				SELECT t.schedule_seat_id
				FROM tickets t
				JOIN payments pay
				ON pay.booking_reference = t.booking_reference
				WHERE pay.transaction_ref=$1
			)
			`,
			txRef,
		)

		if err != nil {
			return fmt.Errorf("releasing seats: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing: %w", err)
	}

	return nil
}