package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	Id int
}

type UserStore interface {
	GetUser(userId int) int
	GetUserGameLibrary(userId int) []int
	GetUserGameAchievementCompletion(userId, gameId int) map[string]int
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

func GetUser(userId int) int {
	return 0
}

func GetUserGameLibrary(userId int) []int {
	return []int{}
}

func GetUserGameAchievementCompletion(userId, gameId int) map[string]int {
	return map[string]int{}
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
