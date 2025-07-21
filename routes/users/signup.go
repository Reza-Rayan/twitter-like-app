package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/dto"
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Signup -> POST method
func Signup(context *gin.Context) {
	var input dto.SignupRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		errors := dto.GetValidationErrors(err)
		context.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password", "error": err.Error()})
		return
	}

	user := models.User{
		Email:    input.Email,
		Username: input.Username,
		Password: hashedPassword,
	}

	if err := user.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}
