package models

import "time"

type Post struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int64     `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Image     *string   `json:"image,omitempty"`
	Likes     []Like    `json:"likes"`
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
