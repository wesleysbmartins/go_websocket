package main

import (
	"go_websocket/internal/infra/server"
	"go_websocket/internal/infra/websocket"
)

func main() {
	hub := websocket.NewHub()
	go hub.Run()

	server.Run()
}
