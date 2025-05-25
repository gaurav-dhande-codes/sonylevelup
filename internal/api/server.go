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

type AchievementLevel struct {
	Name                 string
	OwnedGamesThreshold  int
	AchievementThreshold float64
}

func GetGameAchievementCompletionPercentage(completedAchievement, totalAchievements int) float64 {
	return float64(completedAchievement) / float64(totalAchievements) * 100
}

func (s *SonyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userId := strings.Split(r.URL.String(), "/")[2]
	intUserId, _ := strconv.Atoi(userId)

	userGameLibrary := s.store.GetUserGameLibrary(intUserId)

	// Check if the user owns 10 or fewer games
	if len(userGameLibrary.OwnedGames) <= 10 {
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

	fmt.Fprintf(w, "%s", possibleAchievementLevels[currentIndex].Name)
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
