package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getUserProfile(context *gin.Context) {
	userID := context.GetInt64("userId")

	profile, err := models.GetUserProfile(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user profile", "detail": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User profile fetched successfully",
		"profile": profile,
	})
}
