package api

import (
	"encoding/json"
	"net/http"
)

func GetGameAchievementCompletionPercentage(completedAchievement, totalAchievements int) float64 {
	return float64(completedAchievement) / float64(totalAchievements) * 100
}

func WriteValidResponseToResponseWriter(w http.ResponseWriter, statusCode int, responseBody any) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}
