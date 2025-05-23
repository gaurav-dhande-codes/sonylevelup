package api

import (
	"fmt"
	"net/http"
	"strings"
)

func SonyServer(w http.ResponseWriter, r *http.Request) {
	userId := strings.Split(r.URL.String(), "/")[2]
	fmt.Fprintf(w, GetUserAchievementLevel(userId))
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
