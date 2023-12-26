package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/ivan3177/gohtmlrender/pkg/config"
	"github.com/ivan3177/gohtmlrender/pkg/handlers"
	"github.com/ivan3177/gohtmlrender/pkg/render"
	"log"
	"net/http"
	"time"
)

// appConfig represents the application configuration.
// It stores the following information:
// - UseCache: a boolean indicating whether to use caching or not.
// - Port: an integer representing the port number the application will run on.
// - InProduction: a boolean indicating whether the application is in production or not.
// This configuration is used in various parts of the application to control behavior and settings.
var appConfig = config.AppConfig{
	UseCache:     true,
	Port:         8080,
	InProduction: false,
}

// session is a pointer to a scs.SessionManager struct
// session is used to manage and store session data for a web application
// Example usage:
// Create a new session manager
var session *scs.SessionManager

// main is the entry point of the application.
// It initializes the app configuration, render repository, and handlers repository.
// It sets up the home and about routes.
// Then it listens on the specified port and serves the application.
func main() {
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.InProduction

	renderRepo := render.New(&appConfig)
	handlersRepo := handlers.New(renderRepo, session)

	portNumber := fmt.Sprintf(":%d", appConfig.Port)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(handlersRepo),
	}

	log.Printf("Starting application on port %d", appConfig.Port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Unable to start server on port %d: %v", appConfig.Port, err)
		return
	}

}
