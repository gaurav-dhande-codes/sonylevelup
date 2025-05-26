package api

import "github.com/sonylevelup/internal/model"

// UserStore defines an interface for accessing user-related data.
// Implementations of this interface are responsible for retrieving user
// profiles, game libraries, and achievement completion information.
type UserStore interface {
	// GetUser retrieves a User by their unique userId.
	// Returns the User model or an error if the user is not found or if an internal error occurs.
	GetUser(userId int) (*model.User, error)

	// GetUserGameLibrary retrieves the game library for a given userId.
	// Returns a UserLibrary containing the list of owned games or an error if the user is not found or if an internal error occurs.
	GetUserGameLibrary(userId int) (*model.UserLibrary, error)

	// GetUserGameAchievementCompletion retrieves the achievement completion details for a specific user and game, identified by userId and gameId.
	// Returns a UserGameAchievementCompletion or an error if the user or game is not found or if an internal error occurs.
	GetUserGameAchievementCompletion(userId, gameId int) (*model.UserGameAchievementCompletion, error)
}
