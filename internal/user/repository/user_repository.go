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
	GetUserProfile(id uint) (*models.User, error)
	UpdateUserAvatar(userID uint, avatarURL string) error
	UpdateProfile(u *models.User) error
	FollowUser(f models.Follow) error
	UnfollowUser(userID, unfollowID uint) error
	FindUserByEmail(email string) (*models.User, error)
	SaveOTP(userID uint, otp string, expiresAt time.Time) error
	CheckOTP(userID uint, otp string) (bool, error)
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

// Login -> POST method
func (r *userRepo) Login(email, password string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}
	return &user, nil
}

// GetUserProfile -> GET method
func (r *userRepo) GetUserProfile(id uint) (*models.User, error) {
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUserAvatar -> PATCH method
func (r *userRepo) UpdateUserAvatar(userID uint, avatarURL string) error {
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

// SaveOTP -> POST method
func (r *userRepo) SaveOTP(userID uint, otp string, expiresAt time.Time) error {
	o := models.OTP{UserID: userID, OtpCode: otp, ExpiresAt: expiresAt}
	return db.DB.Create(&o).Error
}

// CheckOTP -> POST method

func (r *userRepo) CheckOTP(userID uint, otp string) (bool, error) {
	var o models.OTP
	if err := db.DB.Where("user_id = ? AND otp_code = ?", userID, otp).
		Order("id desc").First(&o).Error; err != nil {
		return false, err
	}
	if time.Now().After(o.ExpiresAt) {
		return false, nil
	}
	return true, nil
}
