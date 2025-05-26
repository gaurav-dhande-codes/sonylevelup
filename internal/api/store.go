package api

import "github.com/sonylevelup/internal/model"

type UserStore interface {
	GetUser(userId int) (*model.User, error)
	GetUserGameLibrary(userId int) (*model.UserLibrary, error)
	GetUserGameAchievementCompletion(userId, gameId int) (*model.UserGameAchievementCompletion, error)
}
