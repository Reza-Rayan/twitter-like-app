package routes

import (
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// allPosts -> GET method
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

// createPost -> POST method
func createPost(context *gin.Context) {
	userId := context.GetInt64("userId")
	var post models.Post
	if err := context.ShouldBindJSON(&post); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}
	post.UserID = userId
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

// singlePost -> Get method & find by id
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
			"message": "Failed to get post",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Find the post",
		"post":    post,
	})

}

// updatePost -> PUT method & find by id
func updatePost(context *gin.Context) {
	// Checking Authorization
	token := context.GetHeader("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authorization token required",
		})
		return
	}
	// Remove Bearer prefix if present
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	_, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
			"error":   err.Error(),
		})
		return
	}
	// End here

	postId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid post id",
			"error":   err.Error(),
		})
		return
	}
	_, err = models.GetPostByID(postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save post",
			"error":   err.Error(),
		})
		return
	}
	var updatedPost models.Post
	err = context.ShouldBindJSON(&updatedPost)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}
	updatedPost.ID = postId

	err = updatedPost.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save post",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Update Post Was successfully", "post": updatedPost})
}

func deletePost(context *gin.Context) {
	// Checking Authorization
	token := context.GetHeader("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authorization token required",
		})
		return
	}
	// Remove Bearer prefix if present
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	_, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
			"error":   err.Error(),
		})
		return
	}
	// End here

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
			"message": "Failed to get post",
			"error":   err.Error(),
		})
	}
	err = post.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete post",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Delete Post Was successfully", "post": post})
}
