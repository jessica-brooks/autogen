package api

import (
	"github.com/future/automate-similar-text-generation/api/controllers"
)

var server = controllers.Server{}

// Run runs the app
func Run() {
	server.Initialize()
	server.Run(":8080")
}
