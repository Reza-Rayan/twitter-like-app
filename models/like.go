package models

import (
	"database/sql"
	"github.com/Reza-Rayan/twitter-like-app/db"
)

func LikePost(userID, postID int64) error {
	query := `INSERT OR IGNORE INTO likes (user_id, post_id) VALUES (?, ?)`
	_, err := db.DB.Exec(query, userID, postID)
	return err
}

func UnlikePost(userID, postID int64) error {
	query := `DELETE FROM likes WHERE user_id = ? AND post_id = ?`
	_, err := db.DB.Exec(query, userID, postID)
	return err
}

func CountPostLikes(postID int64) (int, error) {
	query := `SELECT COUNT(*) FROM likes WHERE post_id = ?`
	var count int
	err := db.DB.QueryRow(query, postID).Scan(&count)
	return count, err
}

func HasUserLikedPost(userID, postID int64) (bool, error) {
	query := `SELECT 1 FROM likes WHERE user_id = ? AND post_id = ? LIMIT 1`
	var exists int
	err := db.DB.QueryRow(query, userID, postID).Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return true, err
}
