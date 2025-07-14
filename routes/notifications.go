package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getUserNotifications(c *gin.Context) {
	userID := c.GetInt64("userId")

	notifs, err := models.GetUserNotifications(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"notifications": notifs})
}
