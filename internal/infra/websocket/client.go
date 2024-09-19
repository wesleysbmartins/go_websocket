package websocket

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	name string
	room string
	send chan []byte
}

func NewClient(conn *websocket.Conn, name string, room string) *Client {
	return &Client{
		conn: conn,
		name: name,
		room: room,
		send: make(chan []byte, 256),
	}
}

func (c *Client) Read() {
	hub := NewHub()
	defer func() {
		hub.Unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		HandleMessage(c, Message{Room: c.room, Content: message})
	}
}

func (c *Client) Write() {
	ticker := time.NewTicker(10 * time.Second)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.conn.WriteMessage(websocket.TextMessage, message)
		case <-ticker.C:
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
