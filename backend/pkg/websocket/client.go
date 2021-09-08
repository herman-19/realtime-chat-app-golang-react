package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Client struct represents a WebSocket client.
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

// Message struct represents a message.
type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

// Read constantly listens for new messages coming through in this client's
// websocket connection. Messages are be passed to the Pool's broadcast channel,
// which then broadcasts the message to every client in the pool.
func (c *Client) Read() {
	// Defer connection close.
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message received: %+v\n", message)
	}
}