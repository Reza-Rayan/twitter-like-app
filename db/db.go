package db

import (
	"database/sql"
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/config"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	cfg := config.AppConfig.Database

	var err error
	DB, err = sql.Open(cfg.Driver, cfg.Name)
	if err != nil {
		panic(fmt.Sprintf("Could not connect to db: %v", err))
	}

	DB.SetMaxIdleConns(cfg.MaxIdleConns)
	DB.SetMaxOpenConns(cfg.MaxOpenConns)

	createTables()
}

func createTables() {
	// Users table
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	email TEXT NOT NULL UNIQUE,
    	password TEXT NOT NULL,
    	username TEXT,
    	avatar TEXT
	);
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(fmt.Sprintf("Could not create users table: %v", err))
	}

	// Posts table
	createPostsTable := `
	CREATE TABLE IF NOT EXISTS post (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		user_id INTEGER NOT NULL,
		image TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	_, err = DB.Exec(createPostsTable)
	if err != nil {
		panic(fmt.Sprintf("Could not create post table: %v", err))
	}

	// Followers table
	createFollowersTable := `
	CREATE TABLE IF NOT EXISTS follows (
		follower_id INTEGER NOT NULL,
		followee_id INTEGER NOT NULL,
		PRIMARY KEY (follower_id, followee_id),
		FOREIGN KEY (follower_id) REFERENCES users(id),
		FOREIGN KEY (followee_id) REFERENCES users(id)
	);
`
	_, err = DB.Exec(createFollowersTable)
	if err != nil {
		panic(fmt.Sprintf("Could not create followers table: %v", err))
	}

	// notification table
	createNotificationsTable := `
	CREATE TABLE IF NOT EXISTS notifications (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		recipient_id INTEGER NOT NULL,
		sender_id INTEGER NOT NULL,
		type TEXT NOT NULL,
		post_id INTEGER,
		message TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		is_read BOOLEAN DEFAULT 0,
		FOREIGN KEY (recipient_id) REFERENCES users(id),
		FOREIGN KEY (sender_id) REFERENCES users(id),
		FOREIGN KEY (post_id) REFERENCES post(id)
	);
	`
	_, err = DB.Exec(createNotificationsTable)
	if err != nil {
		panic(fmt.Sprintf("Could not create notifications table: %v", err))
	}

	// Like table
	createLikesTable := `
	CREATE TABLE IF NOT EXISTS likes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(user_id, post_id),
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(post_id) REFERENCES post(id)
	);
	`
	_, err = DB.Exec(createLikesTable)
	if err != nil {
		panic(fmt.Sprintf("Could not create likes table: %v", err))
	}

	// OTP table
	createOTPsTable := `
	CREATE TABLE IF NOT EXISTS user_otps (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		otp_code TEXT NOT NULL,
		expires_at DATETIME NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`
	_, err = DB.Exec(createOTPsTable)
	if err != nil {
		panic(fmt.Sprintf("Could not create user_otps table: %v", err))
	}

	createRolesTable := `
	CREATE TABLE IF NOT EXISTS roles (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
     	name TEXT NOT NULL UNIQUE
	);
	`
	_, err = DB.Exec(createRolesTable)
	if err != nil {
		panic(fmt.Sprintf("Could not insert default roles: %v", err))
	}
}
