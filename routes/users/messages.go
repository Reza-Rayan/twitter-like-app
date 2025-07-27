package routes

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Reza-Rayan/twitter-like-app/db"
	"github.com/gin-gonic/gin"
)

func GetMessages(c *gin.Context) {
	userIDInterface, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found"})
		return
	}
	userID := userIDInterface.(int64)

	withIDStr := c.Query("with")
	withID, err := strconv.ParseInt(withIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid receiver id"})
		return
	}

	query := `
	SELECT id, sender_id, receiver_id, content, sent_at FROM messages
	WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)
	ORDER BY sent_at ASC
	`

	rows, err := db.DB.Query(query, userID, withID, withID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	defer rows.Close()

	type Message struct {
		ID         int64  `json:"id"`
		SenderID   int64  `json:"sender_id"`
		ReceiverID int64  `json:"receiver_id"`
		Content    string `json:"content"`
		SentAt     string `json:"sent_at"`
	}

	var messages []Message

	for rows.Next() {
		var m Message
		var sentAtRaw sql.NullString
		err := rows.Scan(&m.ID, &m.SenderID, &m.ReceiverID, &m.Content, &sentAtRaw)
		if err != nil {
			continue
		}
		m.SentAt = sentAtRaw.String
		messages = append(messages, m)
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}
