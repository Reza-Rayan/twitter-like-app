package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/websocket"
	"github.com/gin-gonic/gin"
)

func RegisterChatRoutes(r *gin.RouterGroup) {
	r.GET("/ws", websocket.HandleWebSocket)
}
