package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var UpdateProfile struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

// signup -> POST method
func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse body",
			"error":   err.Error(),
		})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

// login -> POST method
func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse body",
			"error":   err.Error(),
		})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not validate credentials",
			"error":   err.Error(),
		})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not authorize this user",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Logged in successfully",
		"token":   token,
		"user":    user,
	})
}

// updateProfile -> PUT method
func updateUserProfile(context *gin.Context) {
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
