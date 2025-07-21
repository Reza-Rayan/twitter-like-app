package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// DeletePost -> DELETE method & find by id
func DeletePost(context *gin.Context) {

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
	// Clear Cache
	if err := utils.ClearPostsCache(); err != nil {
		log.Printf("Failed to clear posts cache: %v", err)
	}

	context.JSON(http.StatusOK, gin.H{"message": "Delete Post Was successfully", "post": post})
}
