package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/dto"
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Signup -> POST method
func Signup(context *gin.Context) {
	var formRequest dto.SignupRequest
	if err := context.ShouldBindJSON(&formRequest); err != nil {
		errors := dto.GetValidationErrors(err)
		context.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	user := models.User{
		Email:    formRequest.Email,
		Username: formRequest.Username,
		Password: formRequest.Password,
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
