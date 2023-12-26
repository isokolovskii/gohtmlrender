package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// NoSurf is a middleware function that adds Cross-Site Request Forgery (CSRF) protection to the HTTP handler chain.
// It wraps the given handler and returns a new handler that checks for valid CSRF tokens in requests.
func NoSurf(next http.Handler) http.Handler {
	crsfHandler := nosurf.New(next)
	crsfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: !appConfig.InProduction,
		Path:     "/",
		Secure:   appConfig.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return crsfHandler
}

// SessionLoad loads the session for the given HTTP handler.
// The next parameter is the HTTP handler that should be called after the session is loaded.
// It returns an HTTP handler that loads the session and then calls the next handler.
// Example usage:
// ```
// mux := chi.NewRouter()
// ...
// mux.Use(SessionLoad)
// ...
// ```
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
