package post

import "time"

type Post struct {
	ID        int64
	Title     string `binding:"required"`
	Content   string `binding:"required"`
	CreatedAt time.Time
	UserID    int64
	Image     *string `json:"image,omitempty"`
}

type PostWithLikes struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UserID     int64     `json:"user_id"`
	Image      *string   `json:"image,omitempty"`
	LikesCount int       `json:"likes_count"`
}
