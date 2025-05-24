package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/sonylevelup/internal/pkg"
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

	userGameLibrary := s.store.GetUserGameLibrary(intUserId)

	if len(userGameLibrary.OwnedGames) <= 10 {
		fmt.Fprintf(w, pkg.NoAchievementLevel)
		return
	} else {
		fmt.Fprintf(w, pkg.BronzeAchievementLevel)
		return
	}
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
