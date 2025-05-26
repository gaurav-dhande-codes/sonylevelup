package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sonylevelup/internal/model"
	"github.com/sonylevelup/internal/pkg"
)

// GetUserAchievementLevel handles HTTP requests to retrieve the achievement 
// level for a given user. It extracts the userId from the request, fetches the 
// user's data and game library from the store, calculates the achievement level
// based on owned games and achievement completions, and responds with the 
// user's achievement level or an appropriate error.
func (s *SonyServer) GetUserAchievementLevel(w http.ResponseWriter, r *http.Request) {
	// Extract userId from the URL path variables
	vars := mux.Vars(r)
	userId := vars["userId"]

	// Convert userId string to integer
	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		pkg.ErrorLogger.Printf("Failed to parse User ID %q: %v", userId, err)
		pkg.WriteErrorResponseToResponseWriter(w, http.StatusBadRequest, pkg.ErrInvalidUserID)

		return
	}

	// Fetch user data from the data store
	user, err := s.store.GetUser(intUserId)
	if err != nil {
		if err == pkg.ErrUserNotFound {
			pkg.ErrorLogger.Printf("User ID %d does not exist in store: %v", intUserId, err)
			pkg.WriteErrorResponseToResponseWriter(w, http.StatusNotFound, err)
		} else {
			pkg.ErrorLogger.Printf("Failed unexpectedly when fetching user with User ID %d from store: %v", intUserId, err)
			pkg.WriteErrorResponseToResponseWriter(w, http.StatusInternalServerError, pkg.ErrInternalServerError)
		}

		return
	}

	// Initialize UserAchievementLevel response structure with user data
	userAchievementLevel := model.UserAchievementLevel{
		User: user,
	}

	// Fetch the user's game library (list of owned games)
	userGameLibrary, err := s.store.GetUserGameLibrary(intUserId)
	if err != nil {
		if err == pkg.ErrUserNotFound {
			pkg.ErrorLogger.Printf("User ID %d does not exist in store: %v", intUserId, err)
			pkg.WriteErrorResponseToResponseWriter(w, http.StatusNotFound, err)
		} else {
			pkg.ErrorLogger.Printf("Failed unexpectedly when fetching user game library with User ID %d from store: %v", intUserId, err)
			pkg.WriteErrorResponseToResponseWriter(w, http.StatusInternalServerError, pkg.ErrInternalServerError)
		}

		return
	}

	// If user owns 10 or fewer games, set achievement level to "No Achievement" and respond
	if len(userGameLibrary.OwnedGames) <= 10 {
		userAchievementLevel.AchievementLevel = pkg.NoAchievementLevel
		WriteValidResponseToResponseWriter(w, http.StatusOK, userAchievementLevel)

		return
	}

	// Retrieve predefined achievement levels
	achievementLevels := model.GetAchievementLevels()
	currentIndex := 0

	// Iterate over each owned game to determine the user's achievement level
	for _, game := range userGameLibrary.OwnedGames {
		// If current achievement level is Bronze, no need to check further
		if achievementLevels[currentIndex].Name == pkg.BronzeAchievementLevel {
			userAchievementLevel.AchievementLevel = pkg.BronzeAchievementLevel
			WriteValidResponseToResponseWriter(w, http.StatusOK, userAchievementLevel)

			return
		}

		// Fetch the user's achievement completion data for the current game
		gameAchievementCompletion, err := s.store.GetUserGameAchievementCompletion(intUserId, game.Id)
		if err != nil {
			if err == pkg.ErrUserOrGameNotFound {
				pkg.ErrorLogger.Printf("Either User ID %d or Game ID %d does not exist in store: %v", intUserId, game.Id, err)
				pkg.WriteErrorResponseToResponseWriter(w, http.StatusNotFound, err)
			} else {
				pkg.ErrorLogger.Printf("Failed unexpectedly when fetching user game achievement completion with User ID %d and Game ID %d from store: %v", intUserId, game.Id, err)
				pkg.WriteErrorResponseToResponseWriter(w, http.StatusInternalServerError, pkg.ErrInternalServerError)
			}

			return
		}

		// Calculate completion percentage for the current game achievements
		gameAchievementCompletionPercentage := GetGameAchievementCompletionPercentage(
			gameAchievementCompletion.TotalCompletedAchievements,
			gameAchievementCompletion.Game.TotalAvailableAchievements)

		// Update achievement level index based on thresholds and user's progress
		for currentIndex+1 < len(achievementLevels) {
			if gameAchievementCompletionPercentage > achievementLevels[currentIndex].AchievementThreshold &&
				len(userGameLibrary.OwnedGames) > achievementLevels[currentIndex].OwnedGamesThreshold {
				break
			} else {
				currentIndex++
			}
		}
	}

	// Set the final achievement level after evaluating all games
	userAchievementLevel.AchievementLevel = achievementLevels[currentIndex].Name
	WriteValidResponseToResponseWriter(w, http.StatusOK, userAchievementLevel)
}
