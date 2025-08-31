package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

type Client struct {
	ID   int64
	Conn *websocket.Conn
	Send chan []byte
}

type Hub struct {
	Clients    map[int64]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
	mu         sync.Mutex
}

// NewHub  -> Create New HUB if does not exists
func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[int64]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
	}
}

// Run -> event loop runner
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			h.Clients[client.ID] = client
			h.mu.Unlock()
			fmt.Printf("✅ Client %d connected\n", client.ID)

		case client := <-h.Unregister:
			h.mu.Lock()
			if _, ok := h.Clients[client.ID]; ok {
				delete(h.Clients, client.ID)
				close(client.Send)
			}
			h.mu.Unlock()
			fmt.Printf("❌ Client %d disconnected\n", client.ID)

		case message := <-h.Broadcast:
			h.mu.Lock()
			for _, client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client.ID)
				}
			}
			h.mu.Unlock()
		}
	}
}
