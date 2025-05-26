package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sonylevelup/internal/pkg"
)

// SonyServer represents the main HTTP server handling API requests. It holds a
// reference to a UserStore for data operations and a Gorilla mux router for
// routing.
type SonyServer struct {
	store  UserStore
	router *mux.Router
}

// NewSonyServer creates and returns a new SonyServer instance configured with
// the given UserStore. It initializes the HTTP router, sets up the route
// handlers, and applies middleware.
func NewSonyServer(store UserStore) *SonyServer {
	router := mux.NewRouter()
	server := &SonyServer{
		store,
		router,
	}

	// Register route for fetching user achievement level by user ID
	router.HandleFunc("/users/{userId}/achievement-level", server.GetUserAchievementLevel).Methods("GET")

	// Apply logging middleware to all routes
	router.Use(LoggingMiddleware)

	return server
}

// ServeHTTP implements the http.Handler interface, allowing SonyServer to
// handle HTTP requests. It delegates request handling to the underlying Gorilla
// mux router.
func (s *SonyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// LoggingMiddleware is a middleware function that logs incoming HTTP requests
// and their duration. It also recovers from panics during request processing to
// prevent server crashes, returning a 500 Internal Server Error if a panic
// occurs.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		pkg.InfoLogger.Printf("Received request: %s %s from %s", r.Method, r.RequestURI, r.RemoteAddr)

		// Recover from panic to prevent server crash
		defer func() {
			if rec := recover(); rec != nil {
				pkg.ErrorLogger.Printf("Panic recovered: %v", rec)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)

		// Log completion time for the request
		pkg.InfoLogger.Printf("Completed %s %s in %v", r.Method, r.RequestURI, time.Since(start))
	})
}
