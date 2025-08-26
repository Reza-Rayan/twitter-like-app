package service

import (
	"github.com/Reza-Rayan/twitter-like-app/internal/models"
	"github.com/Reza-Rayan/twitter-like-app/internal/user/repository"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"time"
)

type UserService interface {
	Signup(u *models.User) error
	GetUserProfile(int64) (*models.User, error)
	UpdateUserAvatar(userID int64, avatarURL string) error
	UpdateProfile(u *models.User) error
	FollowUser(follow models.Follow) error
	UnfollowUser(userID int64, unfollowID int64) error

	Login(email, password string) (*models.User, error)
	GenerateOTP(email string) (string, error)
	VerifyOTP(email, otp string) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Signup(u *models.User) error {
	return s.repo.Save(u)
}

func (s *userService) Login(email, password string) (*models.User, error) {
	return s.repo.Login(email, password)
}

func (s *userService) GetUserProfile(id int64) (*models.User, error) {
	return s.repo.GetUserProfile(id)
}

func (s *userService) UpdateUserAvatar(userID int64, avatarURL string) error {
	return s.repo.UpdateUserAvatar(userID, avatarURL)
}

func (s *userService) UpdateProfile(u *models.User) error {
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

	// Send OTP to Email in background
	go utils.SendOTPEmail(user.Email, otp)

	return otp, nil
}

func (s *userService) VerifyOTP(email, otp string) (*models.User, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	ok, err := s.repo.CheckOTP(user.ID, otp)
	if err != nil || !ok {
		return nil, err
	}
	return user, nil
}
