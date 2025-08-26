package routes

import (
	"github.com/Reza-Rayan/twitter-like-app/internal/post/handler"
	"github.com/Reza-Rayan/twitter-like-app/internal/post/repository"
	postService "github.com/Reza-Rayan/twitter-like-app/internal/post/service"
	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(router *gin.RouterGroup) {
	repo := repository.NewPostRepository()
	service := postService.NewPostService(repo)
	h := handler.NewPostHandler(service)

	postRouter := router.Group("/posts")
	postRouter.POST("/:id/like", h.LikePost)
	postRouter.DELETE("/:id/like", h.UnlikePost)
}
