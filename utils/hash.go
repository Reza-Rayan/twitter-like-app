package utils

import "golang.org/x/crypto/bcrypt"

// Hashing password by bcrypt for store user password & increase security

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Unhashing password for validate the password is correct or not

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
