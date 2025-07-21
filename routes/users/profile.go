package routes

import (
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserProfile -> GET method
func GetUserProfile(context *gin.Context) {
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

// UpdateAvatar -> PUT method
func UpdateAvatar(context *gin.Context) {
	userID := context.GetInt64("userId")

	file, err := context.FormFile("avatar")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Avatar file is required", "error": err.Error()})
		return
	}

	uploadPath := fmt.Sprintf("uploads/avatars/%d_%s", userID, file.Filename)
	err = context.SaveUploadedFile(file, uploadPath)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save avatar", "error": err.Error()})
		return
	}

	// full URL
	avatarURL := fmt.Sprintf("%s:%d/%s", utils.GetBaseURL(), utils.GetPort(), uploadPath)

	err = models.UpdateUserAvatar(userID, avatarURL)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update avatar", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Avatar updated successfully", "avatar": avatarURL})
}

// UpdateUserProfile -> PUT method
func UpdateUserProfile(context *gin.Context) {
	userId := context.GetInt64("userId")

	// Get form values
	email := context.PostForm("email")
	username := context.PostForm("username")
	password := context.PostForm("password")

	user, err := models.GetUserByID(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "User not found", "error": err.Error()})
		return
	}

	if email != "" {
		user.Email = email
	}
	if username != "" {
		user.Username = username
	}
	if password != "" {
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password", "error": err.Error()})
			return
		}
		user.Password = hashedPassword
	}

	err = user.UpdateProfile()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update profile", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully", "user": user})
}
