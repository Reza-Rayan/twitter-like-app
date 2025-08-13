package routes

import (
	"database/sql"
	"github.com/Reza-Rayan/twitter-like-app/internal/post/handler"
	postRepo "github.com/Reza-Rayan/twitter-like-app/internal/post/repository"
	postService "github.com/Reza-Rayan/twitter-like-app/internal/post/service"
	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(router *gin.RouterGroup, db *sql.DB) {
	repo := postRepo.NewPostRepository(db)
	service := postService.NewPostService(repo)
	h := handler.NewPostHandler(service)

	postRouter := router.Group("/posts")
	postRouter.GET("/", h.GetAllPosts)
	postRouter.GET("/:id", h.GetPostByID)
	postRouter.POST("/", h.CreatePost)
	postRouter.PUT("/:id", h.UpdatePost)
	postRouter.DELETE("/:id", h.DeletePost)

	router.POST("/posts/:id/like", h.LikePost)
	router.DELETE("/posts/:id/like", h.UnLikePost)
}
