package models

import (
	"errors"
	"github.com/Reza-Rayan/twitter-like-app/db"
	"github.com/Reza-Rayan/twitter-like-app/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Username string
	Password string `binding:"required"`
}

// Save -> POST method
func (u User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
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

// ValidateCredentials for login  -> POST method
func (u *User) ValidateCredentials() error {
	query := `
		SELECT id, password FROM users WHERE email = ?
	`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("Invalid password")
	}
	return nil
}

// GetUserByID -> helper
func GetUserByID(id int64) (*User, error) {
	query := `SELECT id, email, COALESCE(username, '') FROM users WHERE id = ?`

	var user User
	err := db.DB.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateProfile -> PUT method
func (u *User) UpdateProfile() error {
	query := `
	UPDATE users
	SET email = ?, username = ?, password = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Email, u.Username, hashedPassword, u.ID)
	return err
}
