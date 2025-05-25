package model

type UserLibrary struct {
	User       User
	OwnedGames []Game
}

type Game struct {
	Id                         int
	Title                      string
	TotalAvailableAchievements int
}

type UserGameAchievementCompletion struct {
	User                       User
	Game                       Game
	TotalCompletedAchievements int
}
