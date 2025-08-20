package handler

import (
	"github.com/Reza-Rayan/twitter-like-app/internal/notification/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NotificationHandler struct {
	service service.NotifyService
}

func NewNotificationHandler(s service.NotifyService) *NotificationHandler {
	return &NotificationHandler{service: s}
}

// GetUserNotifications -> GET method
func (h *NotificationHandler) GetUserNotifications(ctx *gin.Context) {
	userID := ctx.GetInt64("userId")

	notifs, err := h.service.GetUserNotifications(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"notifications": notifs})
}
