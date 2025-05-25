package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id    int
	Name  string
	Email string
}

type Game struct {
	Id                         int
	Title                      string
	TotalAvailableAchievements int
}

type UserLibrary struct {
	User       User
	OwnedGames []Game
}

type UserGameAchievementCompletion struct {
	User                       User
	Game                       Game
	TotalCompletedAchievements int
}

type UserStore interface {
	GetUser(userId int) *User
	GetUserGameLibrary(userId int) *UserLibrary
	GetUserGameAchievementCompletion(userId, gameId int) *UserGameAchievementCompletion
}

type SonyServer struct {
	store UserStore
}

func NewSonyServer(store UserStore) *SonyServer {
	return &SonyServer{store}
}

type AchievementLevel struct {
	Name                 string
	OwnedGamesThreshold  int
	AchievementThreshold float64
}

func GetGameAchievementCompletionPercentage(completedAchievement, totalAchievements int) float64 {
	return float64(completedAchievement) / float64(totalAchievements) * 100
}

func (s *SonyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	router.HandleFunc("/user/{userId}/achievement-level", s.GetUserAchievementLevel)

	router.ServeHTTP(w, r)
}

func GetUser(userId int) *User {
	return &User{}
}

func GetUserGameLibrary(userId int) *UserLibrary {
	return &UserLibrary{}
}

func GetUserGameAchievementCompletion(userId, gameId int) *UserGameAchievementCompletion {
	return &UserGameAchievementCompletion{}
}

func GetUserAchievementLevel(userId int) string {
	if userId == 1 {
		return "Bronze"
	}

	if userId == 2 {
		return "Silver"
	}

	return ""
}
