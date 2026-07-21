package services

import (
	"cinema-ticket-booking/backend/dto"
	"cinema-ticket-booking/backend/models"
	"cinema-ticket-booking/backend/utils"
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type Register struct {
	DB *sql.DB
}

func (r *Register) Signup(c *gin.Context) {
	var req dto.SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	var exists bool
	err := r.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)",
		req.Email,
	).Scan(&exists)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email already exists",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Role:         "user",
	}

	query := `
	INSERT INTO users(name,email,role,password_hash)
	VALUES($1,$2,$3,$4)
	RETURNING id
	`

	err = r.DB.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Role,
		user.PasswordHash,
	).Scan(&user.ID)

	if err != nil {

		// Catch PostgreSQL unique constraint
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				c.JSON(http.StatusConflict, gin.H{
					"error": "Email already exists",
				})
				return
			}
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register user",
		})
		return
	}

	user.PasswordHash = ""

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

func (r *Register) Login(c *gin.Context) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	var user models.User

	query := `
	SELECT id,name,email,password_hash,role
	FROM users
	WHERE email=$1
	`

	err := r.DB.QueryRow(query, req.Email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
	)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid email or password",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token, err := utils.GenerateToken(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	user.PasswordHash = ""

	c.JSON(http.StatusOK, gin.H{
		"message":    "Login successful",
		"token":      token,
		"type":       "Bearer",
		"expires_in": 86400,
		"user":       user,
	})
}