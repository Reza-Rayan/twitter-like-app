package repository

import (
	"database/sql"
	"errors"
	"github.com/Reza-Rayan/twitter-like-app/internal/user"
	"github.com/Reza-Rayan/twitter-like-app/utils"
)

type UserRepository interface {
	Save(u *user.User) error
	Login(u *user.User) error
	GetUserProfile(int64) (*user.User, error)
	UpdateUserAvatar(userID int64, avatarURL string) error
	UpdateProfile(u *user.User) error
	FollowUser(user.Follow) error
	UnfollowUser(userID int64, unfollowID int64) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

// Save -> POST method
func (r *userRepo) Save(u *user.User) error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

// Login -> POST method
func (r *userRepo) Login(u *user.User) error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := r.db.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return err
	}

	if !utils.CheckPasswordHash(u.Password, retrievedPassword) {
		return errors.New("invalid password")
	}
	return nil
}

// GetUserProfile -> GET method
func (r *userRepo) GetUserProfile(id int64) (*user.User, error) {
	query := `SELECT id, email, COALESCE(username, ''), COALESCE(avatar, '') FROM users WHERE id = ?`

	var u user.User
	err := r.db.QueryRow(query, id).Scan(&u.ID, &u.Email, &u.Username, &u.Avatar)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// UpdateUserAvatar -> PATCH method
func (r *userRepo) UpdateUserAvatar(userID int64, avatarURL string) error {
	query := `UPDATE users SET avatar = ? WHERE id = ?`
	_, err := r.db.Exec(query, avatarURL, userID)
	return err
}

// UpdateProfile -> PUT method
func (r *userRepo) UpdateProfile(u *user.User) error {
	query := `
		UPDATE users
		SET email = ?, username = ?, password = ?
		WHERE id = ?
	`
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, u.Email, u.Username, hashedPassword, u.ID)
	return err
}
