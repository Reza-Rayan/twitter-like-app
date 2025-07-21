package routes

import (
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// UpdatePost -> PUT method & find by id
func UpdatePost(context *gin.Context) {
	userId := context.GetInt64("userId")

	postId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid post id", "error": err.Error()})
		return
	}

	post, err := models.GetPostByID(postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get post", "error": err.Error()})
		return
	}

	if post.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to edit this post"})
		return
	}

	// Read updated values from multipart/form-data
	title := context.PostForm("title")
	content := context.PostForm("content")

	if title != "" {
		post.Title = title
	}
	if content != "" {
		post.Content = content
	}

	// Optional: Handle new image upload
	file, err := context.FormFile("image")
	if err == nil {
		uploadRelativePath := fmt.Sprintf("uploads/posts/%d_%s", userId, file.Filename)
		err = context.SaveUploadedFile(file, uploadRelativePath)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save image", "error": err.Error()})
			return
		}

		fullImageURL := fmt.Sprintf("%s:%d/%s", utils.GetBaseURL(), utils.GetPort(), uploadRelativePath)
		post.Image = &fullImageURL
	}

	err = post.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update post", "error": err.Error()})
		return
	}
	// Clear Cache
	if err := utils.ClearPostsCache(); err != nil {
		log.Printf("Failed to clear posts cache: %v", err)
	}
	context.JSON(http.StatusOK, gin.H{"message": "Post updated successfully", "post": post})
}
