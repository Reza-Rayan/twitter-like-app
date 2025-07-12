package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "twitter.db")
	if err != nil {
		panic("Could not connect to db")
	}

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	email TEXT NOT NULL UNIQUE,
    	password TEXT NOT NULL,
    	username TEXT
	);
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(fmt.Sprintf("Could not create users table: %v", err))
	}

	createPostsTable := `
	CREATE TABLE IF NOT EXISTS posts (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    title TEXT NOT NULL,
	    content TEXT NOT NULL,
	    created_at DATETIME NOT NULL,
	    user_id INTEGER NOT NULL,
	    FOREIGN KEY(user_id) REFERENCES users(id) 
	);
	`

	_, err = DB.Exec(createPostsTable)
	if err != nil {
		panic(fmt.Sprintf("Could not create posts table: %v", err))
	}
}
