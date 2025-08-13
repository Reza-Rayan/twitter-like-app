package service

import "github.com/Reza-Rayan/twitter-like-app/internal/user"

func (s *userService) FollowUser(follow user.Follow) error {
	return s.repo.FollowUser(follow)
}

func (s *userService) UnfollowUser(userID int64, unfollowID int64) error {
	return s.repo.UnfollowUser(userID, unfollowID)
}
