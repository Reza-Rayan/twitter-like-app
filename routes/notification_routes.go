package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/internal/notification/handler"
	notificationRepo "github.com/Reza-Rayan/twitter-like-app/internal/notification/repository"
	notificationService "github.com/Reza-Rayan/twitter-like-app/internal/notification/service"
	"github.com/gin-gonic/gin"
)

func RegisterNotificationRoutes(router *gin.RouterGroup) {
	repo := notificationRepo.NewNotificationRepository()
	service := notificationService.NewNotificationService(repo)
	h := handler.NewNotificationHandler(service)

	notifyRouter := router.Group("/notifications")
	notifyRouter.GET("/", h.GetUserNotifications)
	notifyRouter.PATCH("/:id/read", h.MarkAsRead)
}
