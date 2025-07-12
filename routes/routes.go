package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	router := server.Group("/v1")
	// Posts Routes -> v1/post/*
	router.GET("/posts", allPosts)
	router.POST("/posts", createPost)
	router.GET("/posts/:id", singlePost)
	router.PUT("/posts/:id", updatePost)
	router.DELETE("/posts/:id", deletePost)
	//	Users Routes -> v1/register && v1/login
	router.POST("/signup", signup)
	router.POST("/login")
}
