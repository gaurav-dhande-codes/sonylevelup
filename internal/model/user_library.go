package model

type UserLibrary struct {
	User       User   `json:"user"`
	OwnedGames []Game `json:"ownedGames"`
}

type Game struct {
	Id                         int    `json:"id"`
	Title                      string `json:"title"`
	TotalAvailableAchievements int    `json:"totalAvailableAchievements"`
}

type UserGameAchievementCompletion struct {
	User                       User `json:"user"`
	Game                       Game `json:"game"`
	TotalCompletedAchievements int  `json:"totalCompletedAchievements"`
}
