package service

import (
	"github.com/Reza-Rayan/twitter-like-app/internal/user"
	"github.com/Reza-Rayan/twitter-like-app/internal/user/repository"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"time"
)

type UserService interface {
	Signup(u *user.User) error
	GetUserProfile(int64) (*user.User, error)
	UpdateUserAvatar(userID int64, avatarURL string) error
	UpdateProfile(u *user.User) error
	FollowUser(follow user.Follow) error
	UnfollowUser(userID int64, unfollowID int64) error

	Login(u *user.User) error
	GenerateOTP(email string) (string, error)
	//VerifyOTP(email, otp string) (*user.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Signup(u *user.User) error {
	return s.repo.Save(u)
}

func (s *userService) Login(u *user.User) error {
	return s.repo.Login(u)
}

func (s *userService) GetUserProfile(id int64) (*user.User, error) {
	return s.repo.GetUserProfile(id)
}

func (s *userService) UpdateUserAvatar(userID int64, avatarURL string) error {
	return s.repo.UpdateUserAvatar(userID, avatarURL)
}

func (s *userService) UpdateProfile(u *user.User) error {
	return s.repo.UpdateProfile(u)
}

func (s *userService) GenerateOTP(email string) (string, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return "", err
	}

	otp := utils.GenerateOTP(5)
	expiredAt := time.Now().Add(5 * time.Minute) // expire after 5 minute

	if err := s.repo.SaveOTP(user.ID, otp, expiredAt); err != nil {
		return "", err
	}

	go utils.SendOTPEmail(user.Email, otp)

	return otp, nil
}
