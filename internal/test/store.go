package test

import (
	"github.com/sonylevelup/internal/model"
	"github.com/sonylevelup/internal/pkg"
)

type StubUserStore struct {
	users []UserData
}

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
