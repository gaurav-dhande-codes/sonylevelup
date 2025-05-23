package api

import (
	"fmt"
	"net/http"
	"strings"
)

type UserStore interface {
	GetUserAchievementLevel(userId string) string
}

type SonyServer struct {
	store UserStore
}

func NewSonyServer(store UserStore) *SonyServer {
	return &SonyServer{store}
}

func (s *SonyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userId := strings.Split(r.URL.String(), "/")[2]
	fmt.Fprintf(w, s.store.GetUserAchievementLevel(userId))
}

func GetUserAchievementLevel(userId string) string {
	if userId == "1" {
		return "Bronze"
	}

	if userId == "2" {
		return "Silver"
	}

	return ""
}
