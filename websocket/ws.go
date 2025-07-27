package websocket

import (
	"github.com/Reza-Rayan/twitter-like-app/models"
	"github.com/Reza-Rayan/twitter-like-app/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

var connections = make(map[int64]*websocket.Conn)
var mu sync.Mutex

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWebSocket(c *gin.Context) {
	// اعتبارسنجی JWT از QueryParam
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token required"})
		return
	}

	userID, err := utils.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	defer func() {
		mu.Lock()
		delete(connections, userID)
		mu.Unlock()
		conn.Close()
	}()

	mu.Lock()
	connections[userID] = conn
	mu.Unlock()

	for {
		var msg map[string]interface{}
		if err := conn.ReadJSON(&msg); err != nil {
			log.Println("Read error:", err)
			break
		}

		receiverIDFloat, ok := msg["receiver_id"].(float64)
		if !ok {
			log.Println("Invalid receiver_id type")
			continue
		}
		receiverID := int64(receiverIDFloat)

		content, ok := msg["content"].(string)
		if !ok {
			log.Println("Invalid content type")
			continue
		}

		message := models.Message{
			SenderID:   userID,
			ReceiverID: receiverID,
			Content:    content,
			SentAt:     time.Now(),
		}
		err := message.Save()
		if err != nil {
			log.Println("DB save error:", err)
		}

		// ارسال پیام به گیرنده اگر آنلاین بود
		mu.Lock()
		receiverConn, ok := connections[receiverID]
		mu.Unlock()

		if ok {
			if err := receiverConn.WriteJSON(message); err != nil {
				log.Println("Write to receiver error:", err)
			}
		}
	}
}
