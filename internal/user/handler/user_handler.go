package handler

import (
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/dto"
	"github.com/Reza-Rayan/twitter-like-app/internal/user"
	"github.com/Reza-Rayan/twitter-like-app/internal/user/service"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// Signup -> POST method
func (h *UserHandler) Signup(ctx *gin.Context) {
	var formRequest dto.SignupRequest
	if err := ctx.ShouldBindJSON(&formRequest); err != nil {
		errors := dto.GetValidationErrors(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}
	newUser := user.User{
		Email:     formRequest.Email,
		Password:  formRequest.Password,
		Username:  formRequest.Username,
		CreatedAt: time.Now(),
	}
	err := h.service.Signup(&newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register user", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "you signed up successfully", "post": newUser})

}

// Login -> POST method
func (h *UserHandler) Login(ctx *gin.Context) {
	var input dto.LoginRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		errors := dto.GetValidationErrors(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	user := user.User{
		Email:    input.Email,
		Password: input.Password,
	}
	if err := h.service.Login(&user); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials", "error": err.Error()})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logged in successfully",
		"token":   token,
		"user":    user,
	})
}

// GetUserProfile -> POST method
func (h *UserHandler) GetUserProfile(ctx *gin.Context) {
	userID := ctx.GetInt64("userId")

	profile, err := h.service.GetUserProfile(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user profile", "detail": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User profile fetched successfully",
		"profile": profile,
	})

}

// UpdateUserAvatar -> PATCH method
func (h *UserHandler) UpdateUserAvatar(ctx *gin.Context) {
	userID := ctx.GetInt64("userId")

	file, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Avatar file is required", "error": err.Error()})
		return
	}

	uploadPath := fmt.Sprintf("uploads/avatars/%d_%s", userID, file.Filename)
	err = ctx.SaveUploadedFile(file, uploadPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save avatar", "error": err.Error()})
		return
	}

	// full URL
	avatarURL := fmt.Sprintf("%s:%d/%s", utils.GetBaseURL(), utils.GetPort(), uploadPath)

	err = h.service.UpdateUserAvatar(userID, avatarURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update avatar", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Avatar updated successfully", "avatar": avatarURL})
}

// UpdateProfile -> PUT method
func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	// Get form values
	email := ctx.PostForm("email")
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	user, err := h.service.GetUserProfile(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "User not found", "error": err.Error()})
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
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password", "error": err.Error()})
			return
		}
		user.Password = hashedPassword
	}

	if err = h.service.UpdateProfile(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update profile", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully", "user": user})
}

// GenerateOTP -> POST method
func (h *UserHandler) GenerateOTP(ctx *gin.Context) {
	type request struct {
		Email string `json:"email" binding:"required,email"`
	}
	var formRequest request
	if err := ctx.ShouldBindJSON(&formRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.service.GenerateOTP(formRequest.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}
