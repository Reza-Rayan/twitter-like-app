package models

import "time"

type OTP struct {
	ID        int64     `gorm:"primaryKey"`
	UserID    int64     `gorm:"not null"`
	OtpCode   string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
}
