package handler

import (
	"github.com/Reza-Rayan/twitter-like-app/internal/post/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PostHandler struct {
	service service.PostService
}

func NewPostHandler(s service.PostService) *PostHandler {
	return &PostHandler{service: s}
}

// Like a post
func (h *PostHandler) LikePost(ctx *gin.Context) {
	userID := ctx.GetInt64("userId")
	postID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	err = h.service.LikePost(userID, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like post", "detail": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Post liked"})
}

// UnlikePost -> POST method
func (h *PostHandler) UnlikePost(ctx *gin.Context) {
	userID := ctx.GetInt64("userId")
	postID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	err = h.service.UnlikePost(userID, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike post", "detail": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Post unliked"})
}
