package main

import (
	"fmt"
	"github.com/ivan3177/gohtmlrender/pkg/config"
	"github.com/ivan3177/gohtmlrender/pkg/handlers"
	"github.com/ivan3177/gohtmlrender/pkg/render"
	"log"
	"net/http"
)

// main is the entry point of the application.
// It initializes the app configuration, render repository, and handlers repository.
// It sets up the home and about routes.
// Then it listens on the specified port and serves the application.
func main() {
	appConfig := config.AppConfig{
		UseCache: true,
		Port:     8080,
	}

	renderRepo := render.New(&appConfig)
	handlersRepo := handlers.New(renderRepo)

	http.HandleFunc("/", handlersRepo.Home)
	http.HandleFunc("/about", handlersRepo.About)

	log.Println("App listening to port", appConfig.Port)

	_ = http.ListenAndServe(fmt.Sprintf(":%d", appConfig.Port), nil)
}
