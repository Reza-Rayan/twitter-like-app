package repository

import (
	"errors"
	"github.com/Reza-Rayan/twitter-like-app/internal/user"
)

// FollowUser -> POST method
func (r *userRepo) FollowUser(f user.Follow) error {
	if f.FollowerID == f.FolloweeID {
		return errors.New("cannot follow yourself")
	}
	query := `INSERT OR IGNORE INTO follows (follower_id, followee_id) VALUES (?, ?)`
	_, err := r.db.Exec(query, f.FollowerID, f.FolloweeID)
	return err
}

// UnfollowUser -> DELETE method
func (r *userRepo) UnfollowUser(userID int64, unfollowID int64) error {
	query := `DELETE FROM follows WHERE follower_id = ? AND followee_id = ?`
	_, err := r.db.Exec(query, userID, unfollowID)
	return err
}
