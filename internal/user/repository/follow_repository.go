package repository

import (
	"errors"
	"github.com/Reza-Rayan/twitter-like-app/db"
	"github.com/Reza-Rayan/twitter-like-app/internal/models"
)

// FollowUser -> POST method
func (r *userRepo) FollowUser(f models.Follow) error {
	if f.FollowerID == f.FolloweeID {
		return errors.New("cannot follow yourself")
	}
	return db.DB.Create(&f).Error
}

// UnfollowUser -> DELETE method
func (r *userRepo) UnfollowUser(userID, unfollowID uint) error {
	return db.DB.Delete(&models.Follow{}, "follower_id = ? AND followee_id = ?", userID, unfollowID).Error
}
