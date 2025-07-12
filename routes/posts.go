package routes

import (
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func allPosts(context *gin.Context) {
	posts, err := models.GetAllPosts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "fetch the get posts",
			"error":   err.Error(),
		})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Get All Posts", "posts": posts})
}

func createPost(context *gin.Context) {
	var post models.Post
	if err := context.ShouldBindJSON(&post); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}
	post.UserID = 1 // static for now TODO: user _id should be selected after add user management
	fmt.Println(post)
	if err := post.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save post",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Your Post Created", "post": post})
}

func singlePost(context *gin.Context) {
	postId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid post id",
			"error":   err.Error(),
		})
		return
	}
	post, err := models.GetPostByID(postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save post",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Find the post",
		"post":    post,
	})

}
