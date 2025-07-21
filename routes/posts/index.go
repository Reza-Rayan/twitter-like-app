package routes

import (
	"encoding/json"
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AllPosts -> GET method
func AllPosts(context *gin.Context) {
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
