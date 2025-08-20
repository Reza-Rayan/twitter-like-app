package user

import "time"

type User struct {
	ID        int64
	Email     string `binding:"required"`
	Username  string
	Password  string `binding:"required"`
	CreatedAt time.Time
	Avatar    string
	RoleID    int64 `default:"1"`
	Phone     string
}

type Follow struct {
	FollowerID int64 `json:"follower_id"`
	FolloweeID int64 `json:"followee_id"`
}

type PublicUser struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
