package routes

import (
	"errors"
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/dto"
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
)

// CreatePost -> POST method
func CreatePost(c *gin.Context) {
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
		uploadRelativePath := fmt.Sprintf("uploads/post/%d_%s", userId, file.Filename)
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
		log.Printf("Failed to clear post cache: %v", err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created", "post": post})
}
