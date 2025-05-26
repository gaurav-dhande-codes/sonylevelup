package model

type UserAchievementLevel struct {
	User             *User   `json:"user"`
	AchievementLevel string `json:"achievementLevel"`
}
