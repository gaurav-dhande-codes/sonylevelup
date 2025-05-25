package model

import "github.com/sonylevelup/internal/pkg"

type AchievementLevel struct {
	Name                 string
	OwnedGamesThreshold  int
	AchievementThreshold float64
}

func GetAchievementLevels() []AchievementLevel {
	PlatinumAchievement := AchievementLevel{pkg.PlatinumAchievementLevel, 50, 99.99}
	GoldAchievement := AchievementLevel{pkg.GoldAchievementLevel, 25, 80.00}
	SilverAchievement := AchievementLevel{pkg.SilverAchievementLevel, 10, 75.00}
	BronzeAchievement := AchievementLevel{pkg.BronzeAchievementLevel, 10, 00.00}

	return []AchievementLevel{
		PlatinumAchievement,
		GoldAchievement,
		SilverAchievement,
		BronzeAchievement,
	}
}
