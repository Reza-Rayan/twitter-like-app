package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/internal/user/handler"
	userRepo "github.com/Reza-Rayan/twitter-like-app/internal/user/repository"
	userService "github.com/Reza-Rayan/twitter-like-app/internal/user/service"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	repo := userRepo.NewUserRepository()
	service := userService.NewUserService(repo)
	h := handler.NewUserHandler(service)

	router.GET("/profile", h.GetUserProfile)
	router.PATCH("/profile/update-avatar", h.UpdateUserAvatar)
	router.PUT("/profile", h.UpdateProfile)

	//	Follow && Unfollow User
	router.POST("/follow/:id", h.FollowUser)
	router.DELETE("/unfollow/:id", h.UnfollowUser)
}
