package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	// Posts Routes
	router.GET("/posts", allPosts)
	router.POST("/posts", createPost)
	router.GET("/posts/:id", singlePost)
}
