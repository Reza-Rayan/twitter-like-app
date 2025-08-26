package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/internal/user/handler"
	userRepo "github.com/Reza-Rayan/twitter-like-app/internal/user/repository"
	userService "github.com/Reza-Rayan/twitter-like-app/internal/user/service"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {

	repo := userRepo.NewUserRepository()
	service := userService.NewUserService(repo)
	h := handler.NewUserHandler(service)

	router.POST("/signup", h.Signup)
	router.POST("/login", h.Login)
	router.POST("/send-otp", h.GenerateOTP)
	router.POST("/verify-otp", h.VerifyOTP)

}
