package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sonylevelup/internal/pkg"
)

func (s *SonyServer) GetUserAchievementLevel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	intUserId, _ := strconv.Atoi(userId)

	userGameLibrary := s.store.GetUserGameLibrary(intUserId)

	// Check if the user owns 10 or fewer games
	if len(userGameLibrary.OwnedGames) <= 10 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, pkg.NoAchievementLevel)
		return
	}

	PlatinumAchievement := AchievementLevel{pkg.PlatinumAchievementLevel, 50, 99.99}
	GoldAchievement := AchievementLevel{pkg.GoldAchievementLevel, 25, 80.00}
	SilverAchievement := AchievementLevel{pkg.SilverAchievementLevel, 10, 75.00}
	BronzeAchievement := AchievementLevel{pkg.BronzeAchievementLevel, 10, 00.00}

	possibleAchievementLevels := []AchievementLevel{
		PlatinumAchievement,
		GoldAchievement,
		SilverAchievement,
		BronzeAchievement,
	}

	currentIndex := 0

	for _, game := range userGameLibrary.OwnedGames {
		if possibleAchievementLevels[currentIndex].Name == pkg.BronzeAchievementLevel {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, pkg.BronzeAchievementLevel)

			return
		}

		gameAchievementCompletion := s.store.GetUserGameAchievementCompletion(intUserId, game.Id)

		// Calculate the completion percentage for the game
		gameAchievementCompletionPercentage := GetGameAchievementCompletionPercentage(
			gameAchievementCompletion.TotalCompletedAchievements,
			gameAchievementCompletion.Game.TotalAvailableAchievements)

		for currentIndex+1 < len(possibleAchievementLevels) {
			if gameAchievementCompletionPercentage > possibleAchievementLevels[currentIndex].AchievementThreshold &&
				len(userGameLibrary.OwnedGames) > possibleAchievementLevels[currentIndex].OwnedGamesThreshold {
				break
			} else {
				currentIndex++
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", possibleAchievementLevels[currentIndex].Name)
}
