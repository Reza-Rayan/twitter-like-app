package routes

import (
	"database/sql"
	"github.com/Reza-Rayan/twitter-like-app/middlewares"
	notifyRoutes "github.com/Reza-Rayan/twitter-like-app/routes/notify"
	userRoutes "github.com/Reza-Rayan/twitter-like-app/routes/users"
	"github.com/Reza-Rayan/twitter-like-app/websocket"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, db *sql.DB) {
	router := server.Group("/v1")
	authenticated := router.Group("/")
	authenticated.Use(middlewares.AuthMiddleware)

	// Posts Routes -> v1/post/*
	RegisterPostRoutes(authenticated, db)

	//	Users Routes -> v1/register && v1/login && v1/profile/*
	RegisterUserRoutes(authenticated, db)

	//	Follow Users -> v1/follow
	authenticated.POST("/follow/:id", userRoutes.FollowUser)
	authenticated.DELETE("/unfollow/:id", userRoutes.UnfollowUser)

	//	Notifications
	authenticated.GET("/notifications", notifyRoutes.GtUserNotifications)

	//  Send message (ws)
	authenticated.GET("/messages", userRoutes.GetMessages)
	router.GET("/ws", websocket.HandleWebSocket)
}
