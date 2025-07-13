package models

import (
	"errors"
	"github.com/Reza-Rayan/twitter-like-app/db"
)

type Follow struct {
	FollowerID int64 `json:"follower_id"`
	FolloweeID int64 `json:"followee_id"`
}

type PublicUser struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// FollowUser a user
func (f *Follow) FollowUser() error {
	if f.FollowerID == f.FolloweeID {
		return errors.New("cannot follow yourself")
	}
	query := `INSERT OR IGNORE INTO follows (follower_id, followee_id) VALUES (?, ?)`
	_, err := db.DB.Exec(query, f.FollowerID, f.FolloweeID)
	return err
}

// UnfollowUser a user
func UnfollowUser(followerID, followeeID int64) error {
	query := `DELETE FROM follows WHERE follower_id = ? AND followee_id = ?`
	_, err := db.DB.Exec(query, followerID, followeeID)
	return err
}

// GetFollowers a user
func GetFollowers(userID int64) ([]PublicUser, error) {
	query := `
		SELECT u.id, u.email, COALESCE(u.username, '') FROM users u
		JOIN follows f ON f.follower_id = u.id
		WHERE f.followee_id = ?
	`
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []PublicUser
	for rows.Next() {
		var u PublicUser
		if err := rows.Scan(&u.ID, &u.Email, &u.Username); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if users == nil {
		users = []PublicUser{}
	}

	return users, nil
}

// GetFollowing a user
func GetFollowing(userID int64) ([]PublicUser, error) {
	query := `
		SELECT u.id, u.email, COALESCE(u.username, '') FROM users u
		JOIN follows f ON f.followee_id = u.id
		WHERE f.follower_id = ?
	`
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []PublicUser
	for rows.Next() {
		var u PublicUser
		if err := rows.Scan(&u.ID, &u.Email, &u.Username); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if users == nil {
		users = []PublicUser{}
	}

	return users, nil
}
