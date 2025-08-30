package routes

import (
	ws "github.com/Reza-Rayan/twitter-like-app/internal/websocket"
	"github.com/gin-gonic/gin"
)

func RegisterMetricsRoutes(router *gin.RouterGroup, mh *ws.MetricsHub) {
	// ws://localhost:5050/v1/ws/metrics?interval=2s
	router.GET("/ws/metrics", func(c *gin.Context) {
		ws.ServeMetricsWs(mh, c)
	})
}
