package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// This is leave here for simplicity, but in a real-world scenario, you should store this in a secure place
const secretKey = "superSecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"useriD": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
