package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// LikePost -> POST method
func LikePost(context *gin.Context) {
	userID := context.GetInt64("userId")
	postID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}
	if err := models.LikePost(userID, postID); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like post", "data": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Post liked"})
}

// UnLikePost -> DELETE method
func UnLikePost(context *gin.Context) {
	userID := context.GetInt64("userId")
	postID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}
	err = models.UnlikePost(userID, postID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike post", "data": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Post unliked"})
}

func getPostsLike(context *gin.Context) {
	postID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"error": "Invalid post ID"})
		return
	}
	count, err := models.CountPostLikes(postID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch like count"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"post_id": postID, "likes": count})
}
