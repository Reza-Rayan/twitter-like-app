package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/dto"
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
	"time"
)

// allPosts -> GET method
func allPosts(context *gin.Context) {
	limit, offset, page, _ := utils.ParsePagination(context.Request)

	// Cache key (include pagination)
	cacheKey := fmt.Sprintf("posts:limit:%d:offset:%d", limit, offset)

	cached, err := utils.GetCache(cacheKey)
	if err == nil {
		var cachedResponse map[string]interface{}
		if jsonErr := json.Unmarshal([]byte(cached), &cachedResponse); jsonErr == nil {
			context.JSON(http.StatusOK, cachedResponse)
			return
		}
	}

	// Cache miss: Get from DB
	posts, totalCount, err := models.GetAllPosts(limit, offset)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch posts",
			"error":   err.Error(),
		})
		return
	}

	totalPages := (totalCount + limit - 1) / limit
	response := gin.H{
		"message":    "Get All Posts",
		"posts":      posts,
		"limit":      limit,
		"page":       page,
		"totalCount": totalCount,
		"totalPages": totalPages,
	}

	// Set cache for 10 minutes
	bytes, _ := json.Marshal(response)
	_ = utils.SetCache(cacheKey, string(bytes), utils.CacheTime)

	context.JSON(http.StatusOK, response)
}

// createPost -> POST method
func createPost(c *gin.Context) {
	userId := c.GetInt64("userId")

	var input dto.CreatePostRequest
	if err := c.ShouldBind(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make(map[string]string)
			for _, fe := range ve {
				out[fe.Field()] = dto.CustomErrorMessage(fe)
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	title := c.PostForm("title")
	content := c.PostForm("content")

	// handle image file
	file, err := c.FormFile("image")
	var imagePath *string
	if err == nil {
		// Save file locally first
		uploadRelativePath := fmt.Sprintf("uploads/posts/%d_%s", userId, file.Filename)
		err = c.SaveUploadedFile(file, uploadRelativePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save image", "error": err.Error()})
			return
		}

		// Build full URL with base url and port
		fullImageURL := fmt.Sprintf("%s:%d/%s", utils.GetBaseURL(), utils.GetPort(), uploadRelativePath)
		imagePath = &fullImageURL
	}

	post := models.Post{
		Title:     title,
		Content:   content,
		UserID:    userId,
		CreatedAt: time.Now(),
		Image:     imagePath,
	}

	if err := post.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save post", "error": err.Error()})
		return
	}
	followers, err := models.GetFollowers(userId)
	user, err := models.GetUserByID(userId)

	if err == nil {
		for _, follower := range followers {
			notification := models.Notification{
				RecipientID: follower.ID,
				SenderID:    userId,
				Type:        "new_post",
				PostID:      &post.ID,
				Message:     fmt.Sprintf("User %s created a new post", user.Email),
			}
			_ = notification.Save()
		}
	}
	// Clear Cache
	if err := utils.ClearPostsCache(); err != nil {
		log.Printf("Failed to clear posts cache: %v", err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created", "post": post})
}

// singlePost -> Get method & find by id
func singlePost(context *gin.Context) {
	postId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid post id",
			"error":   err.Error(),
		})
		return
	}
	post, err := models.GetPostByID(postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get post",
			"error":   err.Error(),
		})
		return
	}
	// Get post likes
	count, err := models.CountPostLikes(postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get post likes",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message":     "Find the post",
		"post":        post,
		"likes_count": count,
	})

}

// updatePost -> PUT method & find by id
func updatePost(context *gin.Context) {
	userId := context.GetInt64("userId")

	postId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid post id", "error": err.Error()})
		return
	}

	post, err := models.GetPostByID(postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get post", "error": err.Error()})
		return
	}

	if post.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to edit this post"})
		return
	}

	// Read updated values from multipart/form-data
	title := context.PostForm("title")
	content := context.PostForm("content")

	if title != "" {
		post.Title = title
	}
	if content != "" {
		post.Content = content
	}

	// Optional: Handle new image upload
	file, err := context.FormFile("image")
	if err == nil {
		uploadRelativePath := fmt.Sprintf("uploads/posts/%d_%s", userId, file.Filename)
		err = context.SaveUploadedFile(file, uploadRelativePath)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save image", "error": err.Error()})
			return
		}

		fullImageURL := fmt.Sprintf("%s:%d/%s", utils.GetBaseURL(), utils.GetPort(), uploadRelativePath)
		post.Image = &fullImageURL
	}

	err = post.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update post", "error": err.Error()})
		return
	}
	// Clear Cache
	if err := utils.ClearPostsCache(); err != nil {
		log.Printf("Failed to clear posts cache: %v", err)
	}
	context.JSON(http.StatusOK, gin.H{"message": "Post updated successfully", "post": post})
}

// deletePost -> DELETE method & find by id
func deletePost(context *gin.Context) {

	userId := context.GetInt64("userId")

	postId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid post id",
			"error":   err.Error(),
		})
		return
	}

	post, err := models.GetPostByID(postId)

	if post.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to edit this post",
		})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get post",
			"error":   err.Error(),
		})
	}
	err = post.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete post",
			"error":   err.Error(),
		})
		return
	}
	// Clear Cache
	if err := utils.ClearPostsCache(); err != nil {
		log.Printf("Failed to clear posts cache: %v", err)
	}

	context.JSON(http.StatusOK, gin.H{"message": "Delete Post Was successfully", "post": post})
}
