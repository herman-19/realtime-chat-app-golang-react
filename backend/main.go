package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader used to upgrade an HTTP connection to WebSocket connection
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	// Check origin of connection to allow us to make requests
	// from the React dev. server to here.
	// (Temporarily not checking to allow any connection.)
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Defines a reader which listens for new messages sent
// to the Websocket endpoint
func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

// Handler function for WebSocket endpoint.
func wsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// Upgrade HTTP connection to a WebSocket connection.
	// Note: upgrader.Upgrade returns *Conn type used to send and receive messages.
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// Listen (indefinitely) for new messages coming
	// through the WebSocket connection.
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server written in Go")
	})
	http.HandleFunc("/ws", wsHandler)
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}