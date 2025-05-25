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

// CustomNewTestUser creates a test user with a customizable game library.
//
// This function allows generating a user with multiple groups of games,
// where each group can specify its own number of games, available achievements,
// and completed achievements. Useful for more complex test scenarios.
//
// Parameters:
//   - id: Unique identifier for the test user.
//   - name: Name of the user; also used to generate the email address.
//   - gamesMetadata: Variadic list of maps. Each map should include the keys:
//   - "numberOfGames": Total games in this batch.
//   - "numberOfAchievements": Achievements available per game.
//   - "completedAchievements": Achievements completed per game.
//
// Each metadata map represents a batch of games, which will be created in sequence.
// The IDs of the games are auto-incremented and continuous across batches.
//
// Returns:
//
//	A UserData struct containing user info and the complete list of games.
//
// Example:
//
//	user := CustomNewTestUser(1, "bob",
//	  map[string]int{"numberOfGames": 2, "numberOfAchievements": 5, "completedAchievements": 3},
//	  map[string]int{"numberOfGames": 1, "numberOfAchievements": 4, "completedAchievements": 2},
//	)
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

// NewTestUserGameLibrary generates a slice of GamesData with specified parameters.
//
// This function is typically used to simulate a user's game library for testing purposes.
//
// Parameters:
//   - numberOfGames: The number of games to generate.
//   - numberOfAchievements: The total number of achievements available per game.
//   - completedAchievements: The number of achievements completed by the user per game.
//   - startIndex: An offset for generating unique, sequential game IDs.
//
// Returns:
//
//	A slice of GamesData representing the generated game library.
//
// Example:
//
//	games := NewTestUserGameLibrary(3, 10, 5, 0)
//	fmt.Println(games[0].Title) // Output: Game_1
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
