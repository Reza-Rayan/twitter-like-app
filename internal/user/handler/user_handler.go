package handler

import (
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/dto"
	"github.com/Reza-Rayan/twitter-like-app/internal/models"
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
	newUser := models.User{
		Email:     formRequest.Email,
		Password:  formRequest.Password,
		Username:  formRequest.Username,
		CreatedAt: time.Now(),
		RoleID:    1,
	}
	err := h.service.Signup(&newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register user", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "you signed up successfully", "user": newUser})
}

// Login -> POST method
func (h *UserHandler) Login(ctx *gin.Context) {
	var input dto.LoginRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		errors := dto.GetValidationErrors(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	user, err := h.service.Login(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials", "error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID, user.RoleID)
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

// GetUserProfile -> GET method
func (h *UserHandler) GetUserProfile(ctx *gin.Context) {
	userID := int64(ctx.GetInt("userId"))

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
	userID := int64(ctx.GetInt("userId"))

	file, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Avatar file is required", "error": err.Error()})
		return
	}

	uploadPath := fmt.Sprintf("uploads/avatars/%d_%s", userID, file.Filename)
	if err := ctx.SaveUploadedFile(file, uploadPath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save avatar", "error": err.Error()})
		return
	}

	// full URL
	avatarURL := fmt.Sprintf("%s:%d/%s", utils.GetBaseURL(), utils.GetPort(), uploadPath)

	if err := h.service.UpdateUserAvatar(userID, avatarURL); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update avatar", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Avatar updated successfully", "avatar": avatarURL})
}

// UpdateProfile -> PUT method
func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	userID := int64(ctx.GetInt("userId"))

	email := ctx.PostForm("email")
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	user, err := h.service.GetUserProfile(userID)
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

// VerifyOTP -> POST method
func (h *UserHandler) VerifyOTP(ctx *gin.Context) {
	type request struct {
		Email string `json:"email" binding:"required,email"`
		OTP   string `json:"otp" binding:"required"`
	}
	var formInput request
	if err := ctx.ShouldBindJSON(&formInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.VerifyOTP(formInput.Email, formInput.OTP)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired OTP", "error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID, user.RoleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logged in successfully with OTP",
		"token":   token,
		"user":    user,
	})
}
