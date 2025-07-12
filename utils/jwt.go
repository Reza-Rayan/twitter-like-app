package utils

import (
	"errors"
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

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return secretKey, nil
	})

	if err != nil {
		return errors.New("Could not parse token.")
	}
	if !parsedToken.Valid {
		return errors.New("Invalid token!")
	}
	//claims, ok := parsedToken.Claims.(jwt.MapClaims)

	//if !ok {
	//	return errors.New("Invalid token claims")
	//}
	//
	//email := claims["email"].(string)
	//userId := claims["userId"].(int64)
	return nil
}
