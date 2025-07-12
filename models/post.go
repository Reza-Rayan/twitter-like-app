package models

import (
	"github.com/Reza-Rayan/twitter-like-app/db"
	"time"
)

type Post struct {
	ID        int64
	Title     string    `binding:"required"`
	Content   string    `binding:"required"`
	CreatedAt time.Time `binding:"required"`
	UserID    int
}

var posts = []Post{}

// Create New -> POST method
func (p Post) Save() error {
	query := `
		INSERT INTO posts(title, content, created_at, user_id)
		VALUES(?, ?, ?, ?)
		`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(p.Title, p.Content, p.CreatedAt, p.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	p.ID = id

	return err
}
func GetAllPosts() []Post {
	return posts
}
