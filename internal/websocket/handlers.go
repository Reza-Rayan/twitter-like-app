package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/Reza-Rayan/twitter-like-app/db"
	"github.com/Reza-Rayan/twitter-like-app/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWs(hub *Hub, c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, _ := strconv.Atoi(userIDStr)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &Client{
		ID:   int64(userID),
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	hub.Register <- client

	go client.writePump()
	go client.readPump(hub)
}

func (c *Client) readPump(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println("❌ ReadMessage error:", err)
			break
		}

		var msg models.Message
		if err := json.Unmarshal(message, &msg); err != nil {
			fmt.Println("❌ JSON Unmarshal failed:", err, "message:", string(message))
			continue
		}

		msg.SenderID = int64(c.ID)
		msg.CreatedAt = time.Now()

		fmt.Println("💾 Saving message:", msg)

		if err := db.DB.Create(&msg).Error; err != nil {
			fmt.Println("❌ Failed to save message:", err)
		} else {
			fmt.Println("✅ Message saved successfully!")
		}

		hub.Broadcast <- message
	}
}

func (c *Client) writePump() {
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}
