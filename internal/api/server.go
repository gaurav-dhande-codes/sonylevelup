package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

func (s *SonyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userId := strings.Split(r.URL.String(), "/")[2]
	intUserId, _ := strconv.Atoi(userId)

	_ = s.store.GetUser(intUserId)
	_ = s.store.GetUserGameLibrary(intUserId)
	_ = s.store.GetUserGameAchievementCompletion(intUserId, 0)

	fmt.Fprintf(w, "%s", GetUserAchievementLevel(intUserId))
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