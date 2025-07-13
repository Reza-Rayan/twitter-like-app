package models

import (
	"errors"
	"github.com/Reza-Rayan/twitter-like-app/db"
)

type Follow struct {
	FollowerID int64 `json:"follower_id"`
	FolloweeID int64 `json:"followee_id"`
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
