package models

import "time"

type OTP struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	OtpCode   string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
}
