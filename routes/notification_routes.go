package routes

import (
	"database/sql"
	"github.com/Reza-Rayan/twitter-like-app/internal/notification/handler"
	notificationRepo "github.com/Reza-Rayan/twitter-like-app/internal/notification/repository"
	NotificationService "github.com/Reza-Rayan/twitter-like-app/internal/notification/service"
	"github.com/gin-gonic/gin"
)

func RegisterNotificationRoutes(router *gin.RouterGroup, db *sql.DB) {
	repo := notificationRepo.NewNotificationRepository(db)
	service := NotificationService.NewNotificationService(repo)
	h := handler.NewNotificationHandler(service)

	router.GET("/notifications", h.GetUserNotifications)

}
