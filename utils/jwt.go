package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "mySecretKey1234!@#$"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  "",
		"userId": "",
		"exp":    time.Now().Add(time.Hour * 72).Unix(), // 3 days
	})

	return token.SignedString([]byte(secretKey))
}
