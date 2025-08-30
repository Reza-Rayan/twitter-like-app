package service

import (
	"github.com/Reza-Rayan/twitter-like-app/internal/models"
)

func (s *userService) FollowUser(follow models.Follow) error {
	return s.repo.FollowUser(follow)
}

func (s *userService) UnfollowUser(userID int64, unfollowID int64) error {
	return s.repo.UnfollowUser(userID, unfollowID)
}
