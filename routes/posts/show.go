package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// SinglePost -> Get method & find by id
func SinglePost(context *gin.Context) {
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
	// Get post likes
	count, err := models.CountPostLikes(postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get post likes",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message":     "Find the post",
		"post":        post,
		"likes_count": count,
	})

}
