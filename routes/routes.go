package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/internal/websocket"
	"github.com/Reza-Rayan/twitter-like-app/middlewares"
	"github.com/gin-gonic/gin"
	"time"
)

func RegisterRoutes(server *gin.Engine) {
	router := server.Group("/v1")
	authenticated := router.Group("/")
	authenticated.Use(middlewares.AuthMiddleware)

	// Posts Routes -> v1/post/*
	RegisterPostRoutes(authenticated)

	// Auth Routes -> v1/login && v1/signup && v1/send-otp & v1/verify-otp
	RegisterAuthRoutes(router)

	//	Users Routes -> v1/register && v1/login && v1/profile/* && v1/follow && v1/unfollow
	RegisterUserRoutes(authenticated)

	//	Notifications
	RegisterNotificationRoutes(authenticated)

	hub := websocket.NewHub()
	go hub.Run()

	// Message WS Route -> /v1/ws
	RegisterMessageRoutes(authenticated, hub)

	// --- Monitoring Hub ---
	metricsHub := websocket.NewMetricsHub(2 * time.Second) // 2s
	go metricsHub.Run()                                    // Running HUB
	go metricsHub.RunMetricsPublisher()
	RegisterMetricsRoutes(authenticated, metricsHub)

}
