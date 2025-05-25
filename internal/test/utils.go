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
	games := NewTestUserGameLibrary(numberOfGames, numberOfAchievements, completedAchievements, 0)
	return UserData{
		ID:    id,
		Name:  name,
		Email: fmt.Sprintf("%s@sony.com", name),
		Games: games,
	}
}

func CustomNewTestUser(id int, name string, gamesMetadata ...map[string]int) UserData {
	games := []GamesData{}
	startIndex := 0
	for _, gameMetadata := range gamesMetadata {
		customGames := NewTestUserGameLibrary(gameMetadata["numberOfGames"], gameMetadata["numberOfAchievements"], gameMetadata["completedAchievements"], startIndex)
		games = append(games, customGames...)
		startIndex = len(games)
	}

	return UserData{
		ID:    id,
		Name:  name,
		Email: fmt.Sprintf("%s@sony.com", name),
		Games: games,
	}
}

func NewTestUserGameLibrary(numberOfGames, numberOfAchievements, completedAchievements, startIndex int) []GamesData {
	games := []GamesData{}
	for i := 1; i <= numberOfGames; i++ {
		game := GamesData{
			ID:                    startIndex + i,
			Title:                 fmt.Sprintf("Game_%d", startIndex+i),
			AvailableAchievements: numberOfAchievements,
			CompletedAchievements: completedAchievements,
		}
		games = append(games, game)
	}

	return games
}
