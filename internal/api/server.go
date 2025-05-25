package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sonylevelup/internal/model"
)

type SonyServer struct {
	store UserStore
}

type UserStore interface {
	GetUser(userId int) *model.User
	GetUserGameLibrary(userId int) *model.UserLibrary
	GetUserGameAchievementCompletion(userId, gameId int) *model.UserGameAchievementCompletion
}

func NewSonyServer(store UserStore) *SonyServer {
	return &SonyServer{store}
}

func GetGameAchievementCompletionPercentage(completedAchievement, totalAchievements int) float64 {
	return float64(completedAchievement) / float64(totalAchievements) * 100
}

func (s *SonyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	router.HandleFunc("/user/{userId}/achievement-level", s.GetUserAchievementLevel)

	router.ServeHTTP(w, r)
}

func GetUser(userId int) *model.User {
	return &model.User{}
}

func GetUserGameLibrary(userId int) *model.UserLibrary {
	return &model.UserLibrary{}
}

func GetUserGameAchievementCompletion(userId, gameId int) *model.UserGameAchievementCompletion {
	return &model.UserGameAchievementCompletion{}
}
