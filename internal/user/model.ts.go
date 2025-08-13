package user

import "time"

type User struct {
	ID        int64
	Email     string `binding:"required"`
	Username  string
	Password  string `binding:"required"`
	CreatedAt time.Time
	Avatar    string
}
