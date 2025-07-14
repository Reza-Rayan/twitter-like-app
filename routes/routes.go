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
	router.GET("/posts", allPosts)
	router.GET("/posts/:id", singlePost)
	authenticated.POST("/posts", createPost)
	authenticated.PUT("/posts/:id", updatePost)
	authenticated.DELETE("/posts/:id", deletePost)

	//	Users Routes -> v1/register && v1/login
	router.POST("/signup", signup)
	router.POST("/login", login)

	//	Follow Users -> v1/follow
	authenticated.POST("/follow/:id", followUser)
	authenticated.DELETE("/unfollow/:id", unfollowUser)

	//	Profile v1/users
	authenticated.GET("/profile", getUserProfile)
	authenticated.PATCH("/profile/update-avatar", updateAvatar)
}
