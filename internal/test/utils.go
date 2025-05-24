package test

import (
	"fmt"
)

type GamesData struct {
	ID                    int
	Title                 string
	AvailableAchievements int
	CompletedAchievements int
}

type UserData struct {
	ID    int
	Name  string
	Email string
	Games []GamesData
}

func NewTestUser(id int, name string, numberOfGames, numberOfAchievements, completedAchievements int) UserData {
	games := NewTestUserGameLibrary(numberOfGames, numberOfAchievements, completedAchievements)
	return UserData{
		ID:    id,
		Name:  name,
		Email: fmt.Sprintf("%s@sony.com", name),
		Games: games,
	}
}

func NewTestUserGameLibrary(numberOfGames, numberOfAchievements, completedAchievements int) []GamesData {
	games := []GamesData{}
	for i := 1; i <= numberOfGames; i++ {
		game := GamesData{
			ID:                    i,
			Title:                 fmt.Sprintf("Game_%d", i),
			AvailableAchievements: numberOfAchievements,
			CompletedAchievements: completedAchievements,
		}
		games = append(games, game)
	}

	return games
}
