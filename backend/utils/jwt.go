package utils

import (
	"cinema-ticket-booking/backend/models"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

type HasuraClaims struct {
	DefaultRole  string   `json:"x-hasura-default-role"`
	AllowedRoles []string `json:"x-hasura-allowed-roles"`
	UserID       string   `json:"x-hasura-user-id"`
}


type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`

	Hasura HasuraClaims `json:"https://hasura.io/jwt/claims"`

	jwt.RegisteredClaims
}

const tokenTTL = 24 * time.Hour


func GenerateToken(user *models.User) (string, error) {
	if user.Role == "" {
		return "", errors.New("user has no role set")
	}

	expirationTime := time.Now().Add(tokenTTL)

	claims := &Claims{
		Email: user.Email,
		Role:  user.Role,

		Hasura: HasuraClaims{
			DefaultRole:  user.Role,
			AllowedRoles: []string{user.Role},
			UserID:       user.ID.String(),
		},

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-auth-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}


func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}