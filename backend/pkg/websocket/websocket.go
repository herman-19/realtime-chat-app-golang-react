package websocket


import (
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

// Upgrade upgrades HTTP connection to a WebSocket connection.
// Note: upgrader.Upgrade returns *Conn type used to send and receive messages.
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return wsConn, err
	}
	return wsConn, nil
}