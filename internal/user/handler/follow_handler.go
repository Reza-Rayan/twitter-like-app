package handler

import (
	"github.com/Reza-Rayan/twitter-like-app/internal/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FollowUser -> POST method
func (h *UserHandler) FollowUser(ctx *gin.Context) {
	followerID := ctx.GetInt64("userId")
	followeeID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	follow := user.Follow{
		FollowerID: followerID,
		FolloweeID: followeeID,
	}
	if err := h.service.FollowUser(follow); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Followed user"})
}

// UnfollowUser -> DELETE method
func (h *UserHandler) UnfollowUser(ctx *gin.Context) {
	followerID := ctx.GetInt64("userId")
	followeeID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.service.UnfollowUser(followerID, followeeID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Unfollowed user"})

}
