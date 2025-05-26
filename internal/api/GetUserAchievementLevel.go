package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sonylevelup/internal/model"
	"github.com/sonylevelup/internal/pkg"
)

func (s *SonyServer) GetUserAchievementLevel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")

		return
	}

	_, err = s.store.GetUser(intUserId)
	if err != nil {
		if err == pkg.ErrUserNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Not Found")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s", err.Error())
		}

		return
	}

	userGameLibrary, _ := s.store.GetUserGameLibrary(intUserId)

	// Check if the user owns 10 or fewer games
	if len(userGameLibrary.OwnedGames) <= 10 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, pkg.NoAchievementLevel)
		return
	}

	achievementLevels := model.GetAchievementLevels()
	currentIndex := 0

	for _, game := range userGameLibrary.OwnedGames {
		if achievementLevels[currentIndex].Name == pkg.BronzeAchievementLevel {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, pkg.BronzeAchievementLevel)

			return
		}

		gameAchievementCompletion, _ := s.store.GetUserGameAchievementCompletion(intUserId, game.Id)

		// Calculate the completion percentage for the game
		gameAchievementCompletionPercentage := GetGameAchievementCompletionPercentage(
			gameAchievementCompletion.TotalCompletedAchievements,
			gameAchievementCompletion.Game.TotalAvailableAchievements)

		for currentIndex+1 < len(achievementLevels) {
			if gameAchievementCompletionPercentage > achievementLevels[currentIndex].AchievementThreshold &&
				len(userGameLibrary.OwnedGames) > achievementLevels[currentIndex].OwnedGamesThreshold {
				break
			} else {
				currentIndex++
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", achievementLevels[currentIndex].Name)
}
