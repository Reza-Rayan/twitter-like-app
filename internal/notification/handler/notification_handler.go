package handler

import (
	"github.com/Reza-Rayan/twitter-like-app/internal/notification/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

// MarkAsRead -> PATCH method
func (h *NotificationHandler) MarkAsRead(ctx *gin.Context) {
	notificationID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}

	err = h.service.MarkAsRead(notificationID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark notification as read", "detail": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
}
