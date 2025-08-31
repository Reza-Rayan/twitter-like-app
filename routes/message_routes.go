package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/internal/websocket"
	"github.com/gin-gonic/gin"
)

func RegisterMessageRoutes(router *gin.RouterGroup, hub *websocket.Hub) {
	router.GET("/ws", func(c *gin.Context) {
		websocket.ServeWs(hub, c)
	})
}
