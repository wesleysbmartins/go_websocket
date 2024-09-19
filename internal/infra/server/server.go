package server

import (
	"fmt"
	"go_websocket/internal/routes"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()
	router.Use()
	routes.Routes(router)
	http.Handle("/", router)

	corsHandler := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)

	fmt.Println(fmt.Printf("SERVER LISTENNING ON PORT: %v", 8080))
	panic(http.ListenAndServe(fmt.Sprintf(":%v", 8080), corsHandler))
}
