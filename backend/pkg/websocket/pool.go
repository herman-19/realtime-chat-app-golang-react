package websocket

import "fmt"

// Pool of existing connections used
// for concurrent communication.
type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

// CreatePool creates and returns a Pool.
func CreatePool() *Pool {
	p := &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
	return p
}

// Start function constantly listens for clients
// sent to any of the channels of the pool.
func (p *Pool) Start() {
	for {
		// Use select statement to wait on multiple communication operations.
		select {

		// Add new client to pool and send message to all clients.
		case client := <- p.Register:
			p.Clients[client] = true
			fmt.Println("USER ADDED. Pool size: ", len(p.Clients))
			for client := range p.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			}
			break
		
		// Remove client from pool and send message to all clients.
		case client := <- p.Unregister:
			delete(p.Clients, client)
			fmt.Println("USER REMOVED. Pool size: ", len(p.Clients))
			for client := range p.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
			break

		// Broadcasts received message.
		case msg := <- p.Broadcast:
			fmt.Println("Broadcasting message...")
			for client := range p.Clients {
				err := client.Conn.WriteJSON(msg)
				if err != nil {
                    fmt.Println(err)
                    return
                }
			}
			break
		}
	}
}