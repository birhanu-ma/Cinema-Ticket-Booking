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

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Payment struct {
	DB          *sql.DB
	ChapaSecret string
}


// ---------------- Chapa Request / Response ----------------

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


// ---------------- Hasura Action Payload ----------------

type initiatePaymentPayload struct {

	Input struct {
		TicketID uuid.UUID `json:"ticketId"`
	} `json:"input"`

	SessionVariables map[string]string `json:"session_variables"`
}



// ---------------- Initiate Payment Action ----------------

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


	ticketID := payload.Input.TicketID



	// Get user information

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



	// Get ticket price

	var amount string


	err = p.DB.QueryRow(
		`
		SELECT s.price
		FROM tickets t
		JOIN schedules s 
		ON s.id=t.schedule_id
		WHERE t.id=$1
		AND t.user_id=$2
		AND t.status='pending'
		`,
		ticketID,
		userID,
	).Scan(&amount)


	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {

			c.JSON(http.StatusNotFound, gin.H{
				"message":"ticket not found or already paid",
			})

			return
		}


		log.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"failed to get ticket information",
		})

		return
	}



	// Create Chapa transaction reference

	txRef := "cinema-" + uuid.New().String()



	// Save pending payment

	var paymentID uuid.UUID


	err = p.DB.QueryRow(
		`
		INSERT INTO payments
		(
			user_id,
			ticket_id,
			transaction_ref,
			status
		)
		VALUES
		(
			$1,
			$2,
			$3,
			'pending'
		)
		RETURNING id
		`,
		userID,
		ticketID,
		txRef,
	).Scan(&paymentID)



	if err != nil {

		log.Println("payment insert error:", err)


		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}



	// Initialize Chapa payment


	chapaRequest := chapaInitRequest{

		Amount: amount,

		Currency:"ETB",

		Email:email,

		FirstName:firstName,

		LastName:"",

		TxRef:txRef,


		CallbackURL:
        "https://carpentry-wisplike-agonizing.ngrok-free.dev/payment/webhook",


		ReturnURL:
		"http://localhost:3000/payment/success",
	}



	body, err := json.Marshal(chapaRequest)


	if err != nil {

		c.JSON(500, gin.H{
			"message":"failed to create chapa request",
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
			"message":"failed to create request",
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



response, err := http.DefaultClient.Do(request)

if err != nil {
    c.JSON(502, gin.H{
        "message":"failed to connect to Chapa",
    })
    return
}

defer response.Body.Close()


responseBody, err := io.ReadAll(response.Body)

if err != nil {
    c.JSON(500, gin.H{
        "message":"failed reading Chapa response",
    })
    return
}


log.Println("Chapa HTTP status:", response.Status)
log.Println("Chapa response:", string(responseBody))


var chapaResponse chapaInitResponse

err = json.Unmarshal(responseBody, &chapaResponse)

if err != nil {

    c.JSON(500, gin.H{
        "message": "invalid chapa response",
        "response": string(responseBody),
    })

    return
}



	if err != nil {

		c.JSON(500, gin.H{
			"message":"invalid chapa response",
		})

		return
	}



	if chapaResponse.Status!="success" {

		c.JSON(502, gin.H{
			"message":fmt.Sprintf(
				"Chapa error: %s",
				chapaResponse.Message,
			),
		})

		return
	}



	c.JSON(http.StatusOK, gin.H{

		"checkoutUrl":
		chapaResponse.Data.CheckoutURL,

		"txRef":
		txRef,
	})
}





// ---------------- Chapa Webhook ----------------


func (p *Payment) PaymentWebhook(c *gin.Context){


	var payload struct {

		TxRef string `json:"tx_ref"`

		Status string `json:"status"`
	}



	if err:=c.ShouldBindJSON(&payload); err!=nil {

		c.JSON(400,gin.H{
			"error":"invalid webhook",
		})

		return
	}




	verifyRequest,err:=http.NewRequest(
		"GET",
		"https://api.chapa.co/v1/transaction/verify/"+payload.TxRef,
		nil,
	)



	if err!=nil {

		c.JSON(500,gin.H{
			"error":"verification request failed",
		})

		return
	}



	verifyRequest.Header.Set(
		"Authorization",
		"Bearer "+p.ChapaSecret,
	)



	response,err:=http.DefaultClient.Do(verifyRequest)



	if err!=nil {

		c.JSON(500,gin.H{
			"error":"verification failed",
		})

		return
	}


	defer response.Body.Close()



	var verifyResponse struct {

		Data struct {

			Status string `json:"status"`

		} `json:"data"`
	}



	json.NewDecoder(response.Body).
		Decode(&verifyResponse)




	newStatus:="failed"


	if verifyResponse.Data.Status=="success" {

		newStatus="success"
	}




	tx,err:=p.DB.Begin()


	if err!=nil {

		c.JSON(500,gin.H{
			"error":"transaction failed",
		})

		return
	}




	_,err=tx.Exec(
		`
		UPDATE payments
		SET status=$1,
		updated_at=NOW()
		WHERE transaction_ref=$2
		`,
		newStatus,
		payload.TxRef,
	)



	if err!=nil {

		tx.Rollback()

		c.JSON(500,gin.H{
			"error":"payment update failed",
		})

		return
	}




	_,err=tx.Exec(
		`
		UPDATE tickets
		SET status=$1
		WHERE id =
		(
			SELECT ticket_id
			FROM payments
			WHERE transaction_ref=$2
		)
		`,
		newStatus,
		payload.TxRef,
	)



	if err!=nil {

		tx.Rollback()

		c.JSON(500,gin.H{
			"error":"ticket update failed",
		})

		return
	}




	if err:=tx.Commit();err!=nil {

		c.JSON(500,gin.H{
			"error":"commit failed",
		})

		return
	}




	c.JSON(200,gin.H{
		"received":true,
	})
}
