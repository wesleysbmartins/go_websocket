package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"go_websocket/internal/infra/websocket"
	"net/http"
	"time"
)

// @Summary Conexão Websocket
// @Description Estabelece uma conexão WebSocket com a sala especificada
// @Tags websocket
// @Param room query string true "Nome da sala"
// @Param name query string true "Nome do usuário"
// @Success 101 {string} string "Switching Protocols"
// @Router /ws [get]
func WebSocketController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resultErr := make(chan error)

	go websocket.HandleConnection(w, r, resultErr)

	var result error

	for {
		select {
		case result = <-resultErr:
			if result != nil {
				handleError(w, http.StatusInternalServerError, fmt.Sprintf("Failed To Connect to Websocket!\nERROR: %s", result.Error()))
				fmt.Println("Connection Failed!")
			}
			return
		case <-ctx.Done():
			handleError(w, http.StatusRequestTimeout, "Failed To Connect to Websocket!\nTimeout Request")
			fmt.Println("Connection Failed!")
			return
		}
	}
}

func handleError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}
