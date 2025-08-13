package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/dto"
	"github.com/Reza-Rayan/twitter-like-app/internal/post"
	"github.com/Reza-Rayan/twitter-like-app/internal/post/service"
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"time"
)

type PostHandler struct {
	service service.PostService
}

func NewPostHandler(s service.PostService) *PostHandler {
	return &PostHandler{service: s}
}

// CreatePost -> POST method
func (h *PostHandler) CreatePost(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")

	var postRequest dto.CreatePostRequest
	if err := ctx.ShouldBind(&postRequest); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make(map[string]string)
			for _, fe := range ve {
				out[fe.Field()] = dto.CustomErrorMessage(fe)
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	title := ctx.PostForm("title")
	content := ctx.PostForm("content")

	// handle image file
	file, err := ctx.FormFile("image")
	var imagePath *string
	if err == nil {
		// Save file locally first
		uploadRelativePath := fmt.Sprintf("uploads/post/%d_%s", userId, file.Filename)
		err = ctx.SaveUploadedFile(file, uploadRelativePath)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save image", "error": err.Error()})
			return
		}

		// Build full URL with base url and port
		fullImageURL := fmt.Sprintf("%s:%d/%s", utils.GetBaseURL(), utils.GetPort(), uploadRelativePath)
		imagePath = &fullImageURL
	}

	NewPost := post.Post{
		Title:     title,
		Content:   content,
		UserID:    userId,
		CreatedAt: time.Now(),
		Image:     imagePath,
	}
	err = h.service.CreatePost(&NewPost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save post", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Post created", "post": NewPost})

}

// GetAllPosts -> GET method
func (h *PostHandler) GetAllPosts(ctx *gin.Context) {
	limit, offset, page, _ := utils.ParsePagination(ctx.Request)

	// Cache key (include pagination)
	cacheKey := fmt.Sprintf("post:limit:%d:offset:%d", limit, offset)
	cached, err := utils.GetCache(cacheKey)

	if err == nil {
		var cachedResponse map[string]interface{}
		if jsonErr := json.Unmarshal([]byte(cached), &cachedResponse); jsonErr == nil {
			ctx.JSON(http.StatusOK, cachedResponse)
			return
		}
	}
	// Cache miss: Get from DB
	posts, totalCount, err := models.GetAllPosts(limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch post",
			"error":   err.Error(),
		})
		return
	}

	totalPages := (totalCount + limit - 1) / limit
	response := gin.H{
		"message":    "Get All Posts",
		"post":       posts,
		"limit":      limit,
		"page":       page,
		"totalCount": totalCount,
		"totalPages": totalPages,
	}

	// Set cache for 10 minutes
	bytes, _ := json.Marshal(response)
	_ = utils.SetCache(cacheKey, string(bytes), utils.CacheTime)

	ctx.JSON(http.StatusOK, response)
}

// GetPostByID -> GET method
func (h *PostHandler) GetPostByID(ctx *gin.Context) {
	postId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid post id",
			"error":   err.Error(),
		})
		return
	}
	post, err := models.GetPostByID(postId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get post",
			"error":   err.Error(),
		})
		return
	}
	// Get post likes
	count, err := models.CountPostLikes(postId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get post likes",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":     "Find the post",
		"post":        post,
		"likes_count": count,
	})

}

// UpdatePost -> PUT method
func (h *PostHandler) UpdatePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var post post.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "detail": err.Error()})
		return
	}

	post.ID = id
	err = h.service.UpdatePost(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully", "post": post})
}

// DeletePost -> DELETE method
func (h *PostHandler) DeletePost(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	h.service.DeletePost(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted",
	})
}
