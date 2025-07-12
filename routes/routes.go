package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	router := server.Group("/v1")
	// Posts Routes
	router.GET("/posts", allPosts)
	router.POST("/posts", createPost)
	router.GET("/posts/:id", singlePost)
}
