package repository

import (
	"errors"
	"github.com/Reza-Rayan/twitter-like-app/db"
	"github.com/Reza-Rayan/twitter-like-app/internal/models"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"time"
)

type UserRepository interface {
	Save(u *models.User) error
	Login(email, password string) (*models.User, error)
	GetUserProfile(id int64) (*models.User, error)
	UpdateUserAvatar(userID int64, avatarURL string) error
	UpdateProfile(u *models.User) error
	FollowUser(f models.Follow) error
	UnfollowUser(userID, unfollowID int64) error
	FindUserByEmail(email string) (*models.User, error)
	SaveOTP(userID int64, otp string, expiresAt time.Time) error
	CheckOTP(userID int64, otp string) (bool, error)
}

type userRepo struct{}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

// Save -> POST method
func (r *userRepo) Save(u *models.User) error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return db.DB.Create(u).Error
}

// GetUserProfile -> GET method
func (r *userRepo) GetUserProfile(id int64) (*models.User, error) {
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUserAvatar -> PATCH method
func (r *userRepo) UpdateUserAvatar(userID int64, avatarURL string) error {
	return db.DB.Model(&models.User{}).Where("id = ?", userID).Update("avatar", avatarURL).Error
}

// UpdateProfile -> PUT method
func (r *userRepo) UpdateProfile(u *models.User) error {
	if u.Password != "" {
		hashedPassword, _ := utils.HashPassword(u.Password)
		u.Password = hashedPassword
	}
	return db.DB.Save(u).Error
}

// FindUserByEmail -> POST method
func (r *userRepo) FindUserByEmail(email string) (*models.User, error) {
	var u models.User
	if err := db.DB.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
