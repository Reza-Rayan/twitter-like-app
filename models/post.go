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

// Save  New -> POST method
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

// GetAllPosts  -> Get method
func GetAllPosts() ([]Post, error) {
	query := "SELECT * FROM posts"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		rows.Scan(&post.ID, &post.Title, &post.CreatedAt, &post.UserID)
		posts = append(posts, post)
	}
	return posts, nil
}
