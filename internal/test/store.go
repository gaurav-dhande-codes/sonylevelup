package test

import (
	"github.com/sonylevelup/internal/model"
	"github.com/sonylevelup/internal/pkg"
)

type StubUserStore struct {
	users []UserData
}

// GetUser returns a user by their ID from the stubbed user list.
//
// Parameters:
//   - userId: The integer ID of the user to retrieve.
//
// Returns:
//   - A pointer to a model.User if found.
//   - An error if the user is not found.
//
// Example:
//
//	user, err := stubStore.GetUser(1)
//	if err != nil {
//	    t.Fatalf("User not found")
//	}
func (s *StubUserStore) GetUser(userId int) (*model.User, error) {
	for _, testUser := range s.users {
		if testUser.ID == userId {

			return &model.User{
				Id:    testUser.ID,
				Name:  testUser.Name,
				Email: testUser.Email,
			}, nil
		}
	}

	return nil, pkg.ErrUserNotFound
}

// GetUserGameLibrary returns the game library for a user.
//
// Parameters:
//   - userId: The integer ID of the user whose game library is to be retrieved.
//
// Returns:
//   - A pointer to a model.UserLibrary struct containing the user's owned games.
//   - An error if the user is not found.
//
// Example:
//
//	library, err := stubStore.GetUserGameLibrary(1)
//	if err != nil {
//	    t.Errorf("Could not get library: %v", err)
//	}
func (s *StubUserStore) GetUserGameLibrary(userId int) (*model.UserLibrary, error) {
	for _, testUser := range s.users {
		if testUser.ID == userId {
			testOwnedGames := []model.Game{}

			for _, testGame := range testUser.Games {
				testOwnedGames = append(
					testOwnedGames,
					model.Game{
						Id:                         testGame.ID,
						Title:                      testGame.Title,
						TotalAvailableAchievements: testGame.AvailableAchievements,
					})
			}

			return &model.UserLibrary{
				User: model.User{
					Id:    testUser.ID,
					Name:  testUser.Name,
					Email: testUser.Email,
				}, OwnedGames: testOwnedGames}, nil
		}
	}

	return nil, pkg.ErrUserNotFound
}

// GetUserGameAchievementCompletion returns the achievement progress for a user on a specific game.
//
// Parameters:
//   - userId: The user's ID.
//   - gameId: The game ID to retrieve achievement data for.
//
// Returns:
//   - A pointer to model.UserGameAchievementCompletion with detailed user and game achievement data.
//   - An error if either the user or game is not found (pkg.ErrUserNotFound).
//
// Example:
//
//	completion, err := stubStore.GetUserGameAchievementCompletion(1, 101)
//	if err != nil {
//	    t.Logf("Achievement data not found: %v", err)
//	}
func (s *StubUserStore) GetUserGameAchievementCompletion(userId, gameId int) (*model.UserGameAchievementCompletion, error) {
	for _, testUser := range s.users {
		if testUser.ID == userId {

			for _, game := range testUser.Games {
				if game.ID == gameId {

					return &model.UserGameAchievementCompletion{
						User: model.User{
							Id:    testUser.ID,
							Name:  testUser.Name,
							Email: testUser.Email,
						},
						Game: model.Game{
							Id:                         game.ID,
							Title:                      game.Title,
							TotalAvailableAchievements: game.AvailableAchievements,
						},
						TotalCompletedAchievements: game.CompletedAchievements,
					}, nil
				}
			}
		}
	}

	return nil, pkg.ErrUserNotFound
}

// GetAllUsers returns a list of all users in the stub store.
//
// Returns:
//   - A slice of pointers to model.User containing all stub users.
//   - An error if there is a problem retrieving the users (always nil in this stub).
//
// Example:
//
//	users, err := stubStore.GetAllUsers()
//	if err != nil {
//	    t.Fatalf("Failed to get users: %v", err)
//	}
//	for _, user := range users {
//	    t.Logf("User: %v", user)
//	}
func (s *StubUserStore) GetAllUsers() ([]*model.User, error) {
	allUsers := []*model.User{}
	for _, testUser := range s.users {

		allUsers = append(allUsers, &model.User{
			Id:    testUser.ID,
			Name:  testUser.Name,
			Email: testUser.Email,
		})
	}

	return allUsers, nil
}
