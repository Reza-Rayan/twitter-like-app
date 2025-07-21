package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/dto"
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login -> POST method
func Login(context *gin.Context) {
	var input dto.LoginRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		errors := dto.GetValidationErrors(err)
		context.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: input.Password,
	}

	if err := user.ValidateCredentials(); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials", "error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Logged in successfully",
		"token":   token,
		"user":    user,
	})
}
