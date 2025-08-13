package models

import (
	"github.com/Reza-Rayan/twitter-like-app/db"
	"time"
)

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

// Save  New -> POST method
func (p Post) Save() error {
	query := `
		INSERT INTO post(title, content, created_at, user_id, image)
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
func GetAllPosts(limit, offset int) ([]PostWithLikes, int, error) {
	var posts []PostWithLikes

	query := `SELECT * FROM post ORDER BY created_at DESC LIMIT ? OFFSET ?`
	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var post PostWithLikes
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UserID, &post.Image)
		if err != nil {
			return nil, 0, err
		}

		// Fetch like count
		count, err := CountPostLikes(post.ID)
		if err != nil {
			return nil, 0, err
		}
		post.LikesCount = count

		posts = append(posts, post)
	}

	var totalCount int
	countQuery := `SELECT COUNT(*) FROM post`
	err = db.DB.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	return posts, totalCount, nil
}

// GetPostByID -> Get method & find by id
func GetPostByID(id int64) (*Post, error) {
	query := "SELECT * FROM post WHERE id=?"

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
	UPDATE post
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
	query := "DELETE FROM post WHERE id=?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.ID)
	return err
}
