package utils

import (
	"cinema-ticket-booking/backend/models"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))
type Claims struct {
    UserID uuid.UUID `json:"user_id"`
    Email string `json:"email"`
    Role  string `json:"role"`
    jwt.RegisteredClaims
}


func GetEnvOrDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func GenerateToken(user *models.User) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)

    claims := &Claims{
        UserID: user.ID,
        Email: user.Email,
        Role:  user.Role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Issuer:    "go-auth-service",
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(JwtSecret)
}