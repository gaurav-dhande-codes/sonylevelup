package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sonylevelup/internal/pkg"
)

type SonyServer struct {
	store  UserStore
	router *mux.Router
}

func NewSonyServer(store UserStore) *SonyServer {
	router := mux.NewRouter()
	server := &SonyServer{
		store,
		router,
	}

	router.HandleFunc("/users/{userId}/achievement-level", server.GetUserAchievementLevel).Methods("GET")
	router.Use(LoggingMiddleware)
	return server
}

func (s *SonyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		pkg.InfoLogger.Printf("Received request: %s %s from %s", r.Method, r.RequestURI, r.RemoteAddr)

		defer func() {
			if rec := recover(); rec != nil {
				pkg.ErrorLogger.Printf("Panic recovered: %v", rec)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
		pkg.InfoLogger.Printf("Completed %s %s in %v", r.Method, r.RequestURI, time.Since(start))
	})
}
