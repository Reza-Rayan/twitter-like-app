package repository

import (
	"database/sql"
	"github.com/Reza-Rayan/twitter-like-app/internal/post"
	"log"
)

type PostRepository interface {
	Save(p *post.Post) error
	GetAll(limit, offset int) ([]post.PostWithLikes, int, error)
	GetByID(id int64) (*post.Post, error)
	Update(p *post.Post) error
	Delete(id int64) error
	CountLikes(postID int64) (int, error)
	LikePost(userID, postID int64) error
	UnLikePost(userID, postID int64) error
	CountPostLikes(postID int64) (int, error)
}

type postRepo struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepo{db: db}
}

func (r *postRepo) Save(p *post.Post) error {
	query := `INSERT INTO post(title, content, created_at, user_id, image)
		VALUES(?, ?, ?, ?, ?)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(p.Title, p.Content, p.CreatedAt, p.UserID, p.Image)
	if err != nil {
		return err
	}

	p.ID, _ = res.LastInsertId()
	return nil
}

func (r *postRepo) GetAll(limit, offset int) ([]post.PostWithLikes, int, error) {
	query := `SELECT * FROM post ORDER BY created_at DESC LIMIT ? OFFSET ?`
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var posts []post.PostWithLikes
	for rows.Next() {
		var p post.PostWithLikes
		p.LikesCount, _ = r.CountLikes(p.ID)
		if err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.CreatedAt, &p.UserID, &p.Image); err != nil {
			return nil, 0, err
		}
		p.LikesCount, _ = r.CountLikes(p.ID)
		posts = append(posts, p)
	}

	var total int
	r.db.QueryRow(`SELECT COUNT(*) FROM post`).Scan(&total)
	return posts, total, nil
}

func (r *postRepo) GetByID(id int64) (*post.Post, error) {
	query := `SELECT * FROM post WHERE id=?`
	row := r.db.QueryRow(query, id)
	var p post.Post
	err := row.Scan(&p.ID, &p.Title, &p.Content, &p.CreatedAt, &p.UserID, &p.Image)
	log.Println("post", p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *postRepo) Update(p *post.Post) error {
	query := `
	UPDATE post
	SET title=?, content=?, user_id=?, image=?
	WHERE id=?
	`
	_, err := r.db.Exec(query, p.Title, p.Content, p.UserID, p.Image, p.ID)
	return err
}

func (r *postRepo) Delete(id int64) error {
	query := "DELETE FROM post WHERE id=?"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *postRepo) CountLikes(postID int64) (int, error) {
	var count int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM likes WHERE post_id=?`, postID).Scan(&count)
	return count, err
}

func (r *postRepo) LikePost(userID, postID int64) error {
	query := `INSERT OR IGNORE INTO likes (user_id, post_id) VALUES (?, ?)`
	_, err := r.db.Exec(query, userID, postID)
	return err
}

func (r *postRepo) UnLikePost(userID, postID int64) error {
	query := `DELETE FROM likes WHERE user_id = ? AND post_id = ?`
	_, err := r.db.Exec(query, userID, postID)
	return err
}

func (r *postRepo) CountPostLikes(postID int64) (int, error) {
	query := `SELECT COUNT(*) FROM likes WHERE post_id = ?`
	var count int
	err := r.db.QueryRow(query, postID).Scan(&count)
	return count, err
}
