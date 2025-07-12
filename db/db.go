package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to db")
	}

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(5)

	createTables()
}

func createTables() {
	createPostsTable := `
	CREATE TABLE IF NOT EXISTS posts (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    title TEXT NOT NULL,
	    content TEXT NOT NULL,
	    created_at DATETIME NOT NULL,
	    user_id INTEGER NOT NULL
	);
	`

	_, err := DB.Exec(createPostsTable)
	if err != nil {
		panic(fmt.Sprintf("Could not create table: %v", err))
	}
}
