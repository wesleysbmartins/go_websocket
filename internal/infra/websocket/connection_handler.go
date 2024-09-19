package websocket

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnection(w http.ResponseWriter, r *http.Request, resultErr chan error) {
	room := r.URL.Query().Get("room")
	name := r.URL.Query().Get("name")

	if room == "" {
		resultErr <- errors.New("param room is required")
		return
	}

	if name == "" {
		resultErr <- errors.New("param name is required")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		resultErr <- err
		return
	}

	client := NewClient(conn, name, room)

	hub := NewHub()
	hub.Register <- client

	go client.Read()
	go client.Write()

	fmt.Printf("Client %s Registered in Room %s With Success!\n", name, room)

	resultErr <- nil
}
