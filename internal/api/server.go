package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type SonyServer struct {
	store UserStore
}

func NewSonyServer(store UserStore) *SonyServer {
	return &SonyServer{store}
}

func (s *SonyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	router.HandleFunc("/users/{userId}/achievement-level", s.GetUserAchievementLevel)

	router.ServeHTTP(w, r)
}
