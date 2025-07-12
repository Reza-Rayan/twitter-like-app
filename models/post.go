package models

import "time"

type Post struct {
	ID       int
	Title    string    `binding:"required"`
	Content  string    `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID   int
}

var posts = []Post{}

func (p Post) Save() {
	// TODO  add to database later
	posts = append(posts, p)
}

func GetAllPosts() []Post {
	return posts
}
