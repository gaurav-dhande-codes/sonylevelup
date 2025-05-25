package test

import "github.com/sonylevelup/internal/api"

type StubUserStore struct {
	users []UserData
}

func (s *StubUserStore) GetUser(userId int) *api.User {
	for _, testUser := range s.users {
		if testUser.ID == userId {

			return &api.User{
				Id:    testUser.ID,
				Name:  testUser.Name,
				Email: testUser.Email,
			}
		}
	}

	return nil
}

func (s *StubUserStore) GetUserGameLibrary(userId int) *api.UserLibrary {
	for _, testUser := range s.users {
		if testUser.ID == userId {
			testOwnedGames := []api.Game{}

			for _, testGame := range testUser.Games {
				testOwnedGames = append(
					testOwnedGames,
					api.Game{
						Id:                         testGame.ID,
						Title:                      testGame.Title,
						TotalAvailableAchievements: testGame.AvailableAchievements,
					})
			}

			return &api.UserLibrary{
				User: api.User{
					Id:    testUser.ID,
					Name:  testUser.Name,
					Email: testUser.Email,
				}, OwnedGames: testOwnedGames}
		}
	}

	return nil
}

func (s *StubUserStore) GetUserGameAchievementCompletion(userId, gameId int) *api.UserGameAchievementCompletion {
	for _, testUser := range s.users {
		if testUser.ID == userId {

			for _, game := range testUser.Games {
				if game.ID == gameId {

					return &api.UserGameAchievementCompletion{
						User: api.User{
							Id:    testUser.ID,
							Name:  testUser.Name,
							Email: testUser.Email,
						},
						Game: api.Game{
							Id:                         game.ID,
							Title:                      game.Title,
							TotalAvailableAchievements: game.AvailableAchievements,
						},
						TotalCompletedAchievements: game.CompletedAchievements,
					}
				}
			}
		}
	}

	return nil
}
