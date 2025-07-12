package main

import (
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/db"
	"net/http"

	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/posts", allPosts)
	server.POST("/posts", createPost)

	server.Run(":5050")

}

func allPosts(context *gin.Context) {
	posts := models.GetAllPosts()
	context.JSON(http.StatusOK, gin.H{"message": "Get All Posts", "posts": posts})
}

func createPost(context *gin.Context) {
	var post models.Post
	err := context.ShouldBindJSON(&post)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}

	post.ID = 1
	post.UserID = 1
	context.JSON(http.StatusCreated, gin.H{"message": "Your Post Created", "post": post})
}
