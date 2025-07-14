package routes

import (
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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
func createPost(c *gin.Context) {
	userId := c.GetInt64("userId")

	title := c.PostForm("title")
	content := c.PostForm("content")

	// handle image file
	file, err := c.FormFile("image")
	var imagePath *string
	if err == nil {
		// Save file locally first
		uploadRelativePath := fmt.Sprintf("uploads/posts/%d_%s", userId, file.Filename)
		err = c.SaveUploadedFile(file, uploadRelativePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save image", "error": err.Error()})
			return
		}

		// Build full URL with base url and port
		fullImageURL := fmt.Sprintf("%s:%d/%s", utils.GetBaseURL(), utils.GetPort(), uploadRelativePath)
		imagePath = &fullImageURL
	}

	post := models.Post{
		Title:     title,
		Content:   content,
		UserID:    userId,
		CreatedAt: time.Now(),
		Image:     imagePath,
	}

	if err := post.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save post", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created", "post": post})
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
	userId := context.GetInt64("userId")

	postId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid post id",
			"error":   err.Error(),
		})
		return
	}
	post, err := models.GetPostByID(postId)
	if post.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to edit this post",
		})
		return
	}
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

// deletePost -> DELETE method & find by id
func deletePost(context *gin.Context) {

	userId := context.GetInt64("userId")

	postId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid post id",
			"error":   err.Error(),
		})
		return
	}

	post, err := models.GetPostByID(postId)

	if post.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to edit this post",
		})
		return
	}

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
