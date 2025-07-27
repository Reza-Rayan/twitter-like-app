package middlewares

import (
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(context *gin.Context) {
	token := context.GetHeader("Authorization")
	// If no Authorization header, check query param
	if token == "" {
		token = context.Query("token")
	}

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Authorization token required",
		})
		return
	}
	// Remove Bearer prefix if present
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	_, err := utils.VerifyToken(token)

	userID, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
			"error":   err.Error(),
		})
		return
	}
	context.Set("userId", userID)
	context.Next()
}
