package utils

import (
	"cinema-ticket-booking/backend/models"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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

func jwtSecret() ([]byte, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET environment variable is not set")
	}
	return []byte(secret), nil
}

func GenerateToken(user *models.User) (string, error) {
	if user.Role == "" {
		return "", errors.New("user has no role set")
	}

	secret, err := jwtSecret()
	if err != nil {
		return "", err
	}

	now := time.Now()

	claims := &Claims{
		Email: user.Email,
		Role:  user.Role,

		Hasura: HasuraClaims{
			DefaultRole:  user.Role,
			AllowedRoles: []string{user.Role},
			UserID:       user.ID.String(),
		},

		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-auth-service",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(tokenTTL)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func ValidateToken(tokenString string) (*Claims, error) {
	secret, err := jwtSecret()
	if err != nil {
		return nil, err
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}