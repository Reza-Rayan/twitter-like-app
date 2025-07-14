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
	UserID    int64
	Image     *string `json:"image,omitempty"`
}

// Save  New -> POST method
func (p Post) Save() error {
	query := `
		INSERT INTO posts(title, content, created_at, user_id, image)
		VALUES(?, ?, ?, ?, ?)
		`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(p.Title, p.Content, p.CreatedAt, p.UserID, p.Image)
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
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UserID, &post.Image)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// GetPostByID -> Get method & find by id
func GetPostByID(id int64) (*Post, error) {
	query := "SELECT * FROM posts WHERE id=?"

	row := db.DB.QueryRow(query, id)

	var post Post
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UserID, &post.Image)
	if err != nil {
		return nil, err
	}
	return &post, err
}

// Update -> PUT method & find by id
func (post Post) Update() error {
	query := `
	UPDATE posts
	SET title=?, content=?, user_id=?, image=?
	WHERE id=?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.Title, post.Content, post.UserID, post.Image, post.ID)
	return err
}

// Delete -> DELETE method & find by id
func (post Post) Delete() error {
	query := "DELETE FROM posts WHERE id=?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.ID)
	return err
}
