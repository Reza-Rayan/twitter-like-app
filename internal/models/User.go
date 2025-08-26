package models

import "time"

type User struct {
	ID        int64  `gorm:"primaryKey"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Username  string
	Avatar    string
	Phone     string
	RoleID    int64
	Role      Role
	CreatedAt time.Time
}

type Follow struct {
	FollowerID int64 `gorm:"primaryKey"`
	FolloweeID int64 `gorm:"primaryKey"`
}

type PublicUser struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
