package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server contains all the basic objects that helps running the api
type Server struct {
	Router *mux.Router
}

// Initialize initializes the database and routes
func (server *Server) Initialize() {
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

// Run the http server
func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
