package api

import "github.com/sonylevelup/internal/model"

type MockServerUserStore struct{}

func (m *MockServerUserStore) GetUser(userId int) (*model.User, error) {
	return &model.User{}, nil
}

func (m *MockServerUserStore) GetUserGameLibrary(userId int) (*model.UserLibrary, error) {
	return &model.UserLibrary{}, nil
}

func (m *MockServerUserStore) GetUserGameAchievementCompletion(userId, gameId int) (*model.UserGameAchievementCompletion, error) {
	return &model.UserGameAchievementCompletion{}, nil
}
