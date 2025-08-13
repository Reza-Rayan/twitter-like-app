package routes

import (
	"database/sql"
	"github.com/Reza-Rayan/twitter-like-app/internal/user/handler"
	userRepo "github.com/Reza-Rayan/twitter-like-app/internal/user/repository"
	userService "github.com/Reza-Rayan/twitter-like-app/internal/user/service"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup, db *sql.DB) {
	repo := userRepo.NewUserRepository(db)
	service := userService.NewUserService(repo)
	h := handler.NewUserHandler(service)

	router.POST("/signup", h.Signup)
	router.POST("/login", h.Login)
	router.GET("/profile", h.GetUserProfile)
	router.PATCH("/profile/update-avatar", h.UpdateUserAvatar)
	router.PUT("/profile", h.UpdateProfile)
}
