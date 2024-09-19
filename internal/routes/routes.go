package routes

import (
	"fmt"
	_ "go_websocket/docs"
	"go_websocket/internal/controllers"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

// @Summary Mostra uma mensagem de hello
// @Description Retorna uma string "Hello, i am Golang Websocket Server!"
// @Tags health check
// @Success 200 {string} string "Hello, i am Golang Websocket Server!"
// @Router / [get]
func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, i am Golang Websocket Server!")
}

// @title Golang WebSocket
// @version 1.0
// @description Esta é uma aplicação Golang Websocket.
// @termsOfService http://swagger.io/terms/

// @contact.name API Suporte
// @contact.url http://www.swagger.io/support
// @contact.email suporte@swagger.io

// @license.name MIT
// @license.url http://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
func Routes(router *mux.Router) {

	router.HandleFunc("/", healthCheck).Methods("GET")

	router.HandleFunc("/ws", controllers.WebSocketController).Methods("GET")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
