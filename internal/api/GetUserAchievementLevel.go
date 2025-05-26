package api

import (
	"log"
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
		pkg.WriteErrorResponseToResponseWriter(w, http.StatusBadRequest, pkg.ErrInvalidUserID)

		return
	}

	user, err := s.store.GetUser(intUserId)
	if err != nil {
		if err == pkg.ErrUserNotFound {
			pkg.WriteErrorResponseToResponseWriter(w, http.StatusNotFound, err)
		} else {
			pkg.WriteErrorResponseToResponseWriter(w, http.StatusInternalServerError, pkg.ErrInternalServerError)
			log.Println(err)
		}

		return
	}

	userAchievementLevel := model.UserAchievementLevel{
		User: user,
	}

	userGameLibrary, err := s.store.GetUserGameLibrary(intUserId)
	if err != nil {
		if err == pkg.ErrUserNotFound {
			pkg.WriteErrorResponseToResponseWriter(w, http.StatusNotFound, err)
		} else {
			pkg.WriteErrorResponseToResponseWriter(w, http.StatusInternalServerError, pkg.ErrInternalServerError)
			log.Println(err)
		}

		return
	}

	// Check if the user owns 10 or fewer games
	if len(userGameLibrary.OwnedGames) <= 10 {
		userAchievementLevel.AchievementLevel = pkg.NoAchievementLevel
		WriteValidResponseToResponseWriter(w, http.StatusOK, userAchievementLevel)

		return
	}

	achievementLevels := model.GetAchievementLevels()
	currentIndex := 0

	for _, game := range userGameLibrary.OwnedGames {
		if achievementLevels[currentIndex].Name == pkg.BronzeAchievementLevel {
			userAchievementLevel.AchievementLevel = pkg.BronzeAchievementLevel
			WriteValidResponseToResponseWriter(w, http.StatusOK, userAchievementLevel)

			return
		}

		gameAchievementCompletion, err := s.store.GetUserGameAchievementCompletion(intUserId, game.Id)
		if err != nil {
			if err == pkg.ErrUserOrGameNotFound {
				pkg.WriteErrorResponseToResponseWriter(w, http.StatusNotFound, err)
			} else {
				pkg.WriteErrorResponseToResponseWriter(w, http.StatusInternalServerError, pkg.ErrInternalServerError)
				log.Println(err)
			}

			return
		}

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

	userAchievementLevel.AchievementLevel = achievementLevels[currentIndex].Name
	WriteValidResponseToResponseWriter(w, http.StatusOK, userAchievementLevel)
}
