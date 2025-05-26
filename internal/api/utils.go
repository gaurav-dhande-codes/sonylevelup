package api

import (
	"encoding/json"
	"net/http"
)

// GetGameAchievementCompletionPercentage calculates the percentage of
// achievements completed in a game.
//
// Parameters:
//   - completedAchievement: The number of achievements the user has completed.
//   - totalAchievements: The total number of achievements available in the game.
//
// Returns:
//   - A float64 representing the percentage of achievements completed.
//
// Example:
//
//	percentage := GetGameAchievementCompletionPercentage(3, 10) // returns 30.0
func GetGameAchievementCompletionPercentage(completedAchievement, totalAchievements int) float64 {
	return float64(completedAchievement) / float64(totalAchievements) * 100
}

// WriteValidResponseToResponseWriter writes a valid JSON response to the HTTP response writer.
//
// This function sets the HTTP status code and encodes the provided response body as JSON,
// including setting the appropriate "Content-Type" header.
//
// Parameters:
//   - w: The http.ResponseWriter to write the response to.
//   - statusCode: The HTTP status code to set (e.g., http.StatusOK).
//   - responseBody: The response body to encode as JSON. Can be any data structure that is JSON serializable.
//
// Example:
//
//	WriteValidResponseToResponseWriter(w, http.StatusOK, map[string]string{"status": "success"})
func WriteValidResponseToResponseWriter(w http.ResponseWriter, statusCode int, responseBody any) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}
