package websocket

type Message struct {
	Room    string
	Content []byte
}

type MessageJson struct {
	Operation string `json:"operation"`
	Value     any    `json:"Value"`
}
