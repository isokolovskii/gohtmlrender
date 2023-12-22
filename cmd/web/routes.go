package main

import (
	"github.com/bmizerany/pat"
	"github.com/ivan3177/gohtmlrender/pkg/handlers"
	"net/http"
)

// routes is a function that sets up the routing for the application.
// It takes in an instance of the AppConfig struct and a handlers.Repository struct.
// It creates a new pat router and registers the handler functions for the root and about routes.
// It returns the configured router as an http.Handler object.
// The router is responsible for handling incoming HTTP requests and routing them to the appropriate handler functions.
func routes(handlersRepo *handlers.Repository) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlersRepo.Home))
	mux.Get("/about", http.HandlerFunc(handlersRepo.About))

	return mux
}
