package api

import (
	"encoding/json"
	"net/http"
)

// GetGameAchievementCompletionPercentage calculates the percentage of
// achievements completed in a game. It takes the number of completed
// achievements and the total number of available achievements as input, and
// returns the completion percentage as a float64.
func GetGameAchievementCompletionPercentage(completedAchievement, totalAchievements int) float64 {
	return float64(completedAchievement) / float64(totalAchievements) * 100
}

// WriteValidResponseToResponseWriter writes a JSON-encoded responseBody to the
// provided http.ResponseWriter with the given statusCode and sets appropriate
// headers
func WriteValidResponseToResponseWriter(w http.ResponseWriter, statusCode int, responseBody any) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}
