package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ivan3177/gohtmlrender/pkg/handlers"
	"net/http"
)

// routes initializes and configures the Router for the application's routes.
// It sets up middleware for error recovery, CSRF protection, and session management.
// It also registers the routes for the Home and About pages.
// The handler functions for the routes are provided through the handlers.Repository parameter.
// The configured Router is returned as an http.Handler.
//
// Example usage:
//
//	renderRepo := render.New(&appConfig)
//	handlersRepo := handlers.New(renderRepo, session)
//
//	srv := &http.Server{
//	    Addr:    portNumber,
//	    Handler: routes(handlersRepo),
//	}
//	err := srv.ListenAndServe()
//	if err != nil {
//	    log.Fatalf("Unable to start server on port %d: %v", appConfig.Port, err)
//	    return
//	}
func routes(handlersRepo *handlers.Repository) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlersRepo.Home)
	mux.Get("/about", handlersRepo.About)

	return mux
}
