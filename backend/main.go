package main

import (
	"fmt"
	"net/http"

	"github.com/herman-19/golang-realtime-chat-app/pkg/websocket"
)


// wsHandler is the handler function for the WebSocket endpoint.
func wsHandler(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to a WebSocket connection.
	// Note: upgrader.Upgrade returns *Conn type used to send and receive messages.
	wsConn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	// Create new client
	client := &websocket.Client{
		Conn: wsConn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.CreatePool()
	go pool.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server written in Go")
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsHandler(pool, w, r)
	})
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}