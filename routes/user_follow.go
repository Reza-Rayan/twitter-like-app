package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func followUser(context *gin.Context) {
	// Get users IDs
	followerID := context.GetInt64("userId")
	followeeID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	follow := models.Follow{
		FollowerID: followerID,
		FolloweeID: followeeID,
	}

	if err := follow.FollowUser(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Followed user"})
}

func unfollowUser(context *gin.Context) {
	// Get users IDs
	followerID := context.GetInt64("userId")
	followeeID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	err = models.UnfollowUser(followerID, followeeID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Unfollowed user"})

}
