package repository

import (
	"errors"
	"github.com/Reza-Rayan/twitter-like-app/db"
	"github.com/Reza-Rayan/twitter-like-app/internal/models"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"time"
)

// SaveOTP -> POST method
func (r *userRepo) SaveOTP(userID int64, otp string, expiresAt time.Time) error {
	o := models.OTP{UserID: userID, OtpCode: otp, ExpiresAt: expiresAt}
	return db.DB.Create(&o).Error
}

// CheckOTP -> POST method
func (r *userRepo) CheckOTP(userID int64, otp string) (bool, error) {
	var o models.OTP
	if err := db.DB.Where("user_id = ? AND otp_code = ?", userID, otp).
		Order("id desc").First(&o).Error; err != nil {
		return false, err
	}
	if time.Now().After(o.ExpiresAt) {
		return false, nil
	}
	return true, nil
}

// Login -> POST method
func (r *userRepo) Login(email, password string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}
	return &user, nil
}
