package api

import (
	"net/http"

	"github.com/gorilla/mux"
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
	return server
}

func (s *SonyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
