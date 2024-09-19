package websocket

import (
	"fmt"
	"sync"
)

type Hub struct {
	Clients    map[*Client]bool
	Rooms      map[string]map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
}

var instance *Hub
var once sync.Once

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true

			if h.Rooms[client.room] == nil {
				h.Rooms[client.room] = make(map[*Client]bool)
			}
			h.Rooms[client.room][client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				delete(h.Rooms[client.room], client)
				close(client.send)
				client.conn.Close()
				fmt.Printf("Client %s Unregistered!\n", client.name)
			}
		case message := <-h.Broadcast:
			for client := range h.Rooms[message.Room] {
				client.send <- message.Content
				fmt.Printf("Send Message to Client %s.\n", client.name)
			}
		}
	}
}

func NewHub() *Hub {
	once.Do(func() {
		instance = &Hub{
			Clients:    make(map[*Client]bool),
			Rooms:      make(map[string]map[*Client]bool),
			Register:   make(chan *Client),
			Unregister: make(chan *Client),
			Broadcast:  make(chan Message),
		}
	})
	return instance
}
