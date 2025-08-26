package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/middlewares"
	"github.com/gin-gonic/gin"
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

}
