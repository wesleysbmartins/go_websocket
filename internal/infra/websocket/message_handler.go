package websocket

import (
	"encoding/json"
)

func HandleMessage(client *Client, message Message) {
	messageJson := MessageJson{}
	json.Unmarshal(message.Content, &messageJson)

	hub := NewHub()

	if messageJson.Operation == "default" {
		hub.Broadcast <- message
	} else if messageJson.Operation == "unregister" {
		hub.Unregister <- client
	}
}
