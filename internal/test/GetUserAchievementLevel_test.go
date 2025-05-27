package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sonylevelup/internal/api"
	"github.com/sonylevelup/internal/pkg"
)

func TestGetUserAchievementLevelForNoAchievementLevelUsers(t *testing.T) {
	testUsers := []UserData{
		// No Achievement Level Users:
		// 10 or fewer games owned

		// User owns 0 games, basically a new account
		CustomNewTestUser(1, "Garry", map[string]int{
			"numberOfGames":         0,
			"numberOfAchievements":  0,
			"completedAchievements": 0,
		}),

		// User owns 10 games, and has completed 100% achievements in all games.
		CustomNewTestUser(2, "Tom", map[string]int{
			"numberOfGames":         10,
			"numberOfAchievements":  100,
			"completedAchievements": 100,
		}),

		// User owns 10 Games and has completed atleast 90% achievements in all games.
		CustomNewTestUser(3, "Bob", map[string]int{
			"numberOfGames":         10,
			"numberOfAchievements":  100,
			"completedAchievements": 90,
		}),

		// User owns 10 Games and has completed atleast 80% achievements in all games.
		CustomNewTestUser(4, "Luna", map[string]int{
			"numberOfGames":         10,
			"numberOfAchievements":  100,
			"completedAchievements": 80,
		}),

		// User owns 10 Games and has completed atleast 75% achievements in all games.
		CustomNewTestUser(5, "Jerry", map[string]int{
			"numberOfGames":         10,
			"numberOfAchievements":  100,
			"completedAchievements": 75,
		}),
	}
	testStore := StubUserStore{testUsers}
	testServer := api.NewSonyServer(&testStore)

	for _, user := range testUsers {
		t.Run(fmt.Sprintf("Test User %s", user.Name), func(t *testing.T) {
			response := httptest.NewRecorder()
			testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, fmt.Sprint(user.ID)))

			wantedResponse, gotResponse := getExpectedAndReceivedUserAchievementValidResponse(t, user, response, pkg.NoAchievementLevel)

			assertHttpResponseStatus(t, response.Code, http.StatusOK)
			assertContentType(t, response, pkg.JsonContentType)
			assertResponseBody(t, gotResponse, wantedResponse)
		})
	}
}

func TestGetUserAchievementLevelForBronzeAchievementLevelUsers(t *testing.T) {
	testUsers := []UserData{
		// Bronze Achievement Level Users:
		// More than 10 games owned

		// User owns 11 Games and has completed 76% achievements in 10 games and 75% achievements in 1 game.
		CustomNewTestUser(1, "Garry", map[string]int{
			"numberOfGames":         10,
			"numberOfAchievements":  100,
			"completedAchievements": 76,
		}, map[string]int{
			"numberOfGames":         1,
			"numberOfAchievements":  100,
			"completedAchievements": 75,
		}),

		// User owns 26 Games and has completed atleast 81% achievements in 25 games and 75% achievements in 1 game.
		CustomNewTestUser(2, "Tom", map[string]int{
			"numberOfGames":         25,
			"numberOfAchievements":  100,
			"completedAchievements": 81,
		}, map[string]int{
			"numberOfGames":         1,
			"numberOfAchievements":  100,
			"completedAchievements": 75,
		}),

		// User owns 51 Games and has completed atleast 100% achievements in 50 games and 75% achievements in 1 game.
		CustomNewTestUser(3, "Bob", map[string]int{
			"numberOfGames":         50,
			"numberOfAchievements":  100,
			"completedAchievements": 100,
		}, map[string]int{
			"numberOfGames":         1,
			"numberOfAchievements":  100,
			"completedAchievements": 75,
		}),

		// User owns 11 Games and has completed 75% achievements in all games.
		CustomNewTestUser(4, "Luna", map[string]int{
			"numberOfGames":         11,
			"numberOfAchievements":  100,
			"completedAchievements": 75,
		}),

		// User owns 26 Games and has completed 75% achievements in all games.
		CustomNewTestUser(5, "Jerry", map[string]int{
			"numberOfGames":         26,
			"numberOfAchievements":  100,
			"completedAchievements": 75,
		}),

		// User owns 51 Games and has completed 75% achievements in all games.
		CustomNewTestUser(6, "Sally", map[string]int{
			"numberOfGames":         51,
			"numberOfAchievements":  100,
			"completedAchievements": 75,
		}),

		// User owns 11 Games and has completed 0% achievements in all games.
		CustomNewTestUser(7, "Cody", map[string]int{
			"numberOfGames":         11,
			"numberOfAchievements":  100,
			"completedAchievements": 0,
		}),
	}
	testStore := StubUserStore{testUsers}
	testServer := api.NewSonyServer(&testStore)

	for _, user := range testUsers {
		t.Run(fmt.Sprintf("Test User %s", user.Name), func(t *testing.T) {
			response := httptest.NewRecorder()
			testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, fmt.Sprint(user.ID)))

			wantedResponse, gotResponse := getExpectedAndReceivedUserAchievementValidResponse(t, user, response, pkg.BronzeAchievementLevel)

			assertHttpResponseStatus(t, response.Code, http.StatusOK)
			assertContentType(t, response, pkg.JsonContentType)
			assertResponseBody(t, gotResponse, wantedResponse)
		})
	}
}

func TestGetUserAchievementLevelForSilverAchievementLevelUsers(t *testing.T) {
	testUsers := []UserData{
		// Silver Achievement Level Users:
		// Owns more than 10 games and has 75%+ achievements in each

		// User owns 26 Games and has completed 81% achievements in 25 games and 80% achievements in 1 game.
		CustomNewTestUser(1, "Garry", map[string]int{
			"numberOfGames":         25,
			"numberOfAchievements":  100,
			"completedAchievements": 81,
		}, map[string]int{
			"numberOfGames":         1,
			"numberOfAchievements":  100,
			"completedAchievements": 80,
		}),

		// User owns 26 Games and has completed 81% achievements in 25 games and 76% achievements in 1 game.
		CustomNewTestUser(2, "Tom", map[string]int{
			"numberOfGames":         25,
			"numberOfAchievements":  100,
			"completedAchievements": 81,
		}, map[string]int{
			"numberOfGames":         1,
			"numberOfAchievements":  100,
			"completedAchievements": 76,
		}),

		// User owns 51 Games and has completed 100% achievements in 50 games and 80% achievements in 1 game.
		CustomNewTestUser(3, "Bob", map[string]int{
			"numberOfGames":         50,
			"numberOfAchievements":  100,
			"completedAchievements": 100,
		}, map[string]int{
			"numberOfGames":         1,
			"numberOfAchievements":  100,
			"completedAchievements": 80,
		}),

		// User owns 51 Games and has completed 100% achievements in 50 games and 76% achievements in 1 game.
		CustomNewTestUser(4, "Luna", map[string]int{
			"numberOfGames":         50,
			"numberOfAchievements":  100,
			"completedAchievements": 100,
		}, map[string]int{
			"numberOfGames":         1,
			"numberOfAchievements":  100,
			"completedAchievements": 76,
		}),

		// User owns 26 Games and has completed 80% achievements in all games.
		CustomNewTestUser(5, "Jerry", map[string]int{
			"numberOfGames":         26,
			"numberOfAchievements":  100,
			"completedAchievements": 80,
		}),

		// User owns 51 Games and has completed 80% achievements in all games.
		CustomNewTestUser(6, "Sally", map[string]int{
			"numberOfGames":         51,
			"numberOfAchievements":  100,
			"completedAchievements": 80,
		}),

		// User owns 26 Games and has completed 76% achievements in all games.
		CustomNewTestUser(7, "Cody", map[string]int{
			"numberOfGames":         26,
			"numberOfAchievements":  100,
			"completedAchievements": 76,
		}),

		// User owns 51 Games and has completed 76% achievements in all games.
		CustomNewTestUser(8, "Ezra", map[string]int{
			"numberOfGames":         51,
			"numberOfAchievements":  100,
			"completedAchievements": 76,
		}),

		// User owns 11 Games and has completed 76% achievements in all games.
		CustomNewTestUser(9, "Daniel", map[string]int{
			"numberOfGames":         11,
			"numberOfAchievements":  100,
			"completedAchievements": 76,
		}),
	}
	testStore := StubUserStore{testUsers}
	testServer := api.NewSonyServer(&testStore)

	for _, user := range testUsers {
		t.Run(fmt.Sprintf("Test User %s", user.Name), func(t *testing.T) {
			response := httptest.NewRecorder()
			testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, fmt.Sprint(user.ID)))

			wantedResponse, gotResponse := getExpectedAndReceivedUserAchievementValidResponse(t, user, response, pkg.SilverAchievementLevel)

			assertHttpResponseStatus(t, response.Code, http.StatusOK)
			assertContentType(t, response, pkg.JsonContentType)
			assertResponseBody(t, gotResponse, wantedResponse)
		})
	}
}

func TestGetUserAchievementLevelForGoldAchievementLevelUsers(t *testing.T) {
	testUsers := []UserData{
		// Gold Achievement Level Users:
		// Owns 25+ games and has 80%+ achievements in each

		// User owns 51 Games and has completed 100% achievements in 50 games and 99% achievements in 1 game.
		CustomNewTestUser(1, "Garry", map[string]int{
			"numberOfGames":         50,
			"numberOfAchievements":  100,
			"completedAchievements": 100,
		}, map[string]int{
			"numberOfGames":         1,
			"numberOfAchievements":  100,
			"completedAchievements": 99,
		}),

		// User owns 51 Games and has completed 100% achievements in 50 games and 81% achievements in 1 game.
		CustomNewTestUser(2, "Tom", map[string]int{
			"numberOfGames":         50,
			"numberOfAchievements":  100,
			"completedAchievements": 100,
		}, map[string]int{
			"numberOfGames":         1,
			"numberOfAchievements":  100,
			"completedAchievements": 81,
		}),

		// User owns 50 Games and has completed 100% achievements in all games.
		CustomNewTestUser(3, "Bob", map[string]int{
			"numberOfGames":         50,
			"numberOfAchievements":  100,
			"completedAchievements": 100,
		}),

		// User owns 50 Games and has completed 81% achievements in all games.
		CustomNewTestUser(4, "Luna", map[string]int{
			"numberOfGames":         50,
			"numberOfAchievements":  100,
			"completedAchievements": 81,
		}),

		// User owns 26 Games and has completed 81% achievements in all games.
		CustomNewTestUser(5, "Jerry", map[string]int{
			"numberOfGames":         26,
			"numberOfAchievements":  100,
			"completedAchievements": 81,
		}),

		// User owns 26 Games and has completed 100% achievements in all games.
		CustomNewTestUser(6, "Sally", map[string]int{
			"numberOfGames":         26,
			"numberOfAchievements":  100,
			"completedAchievements": 100,
		}),
	}
	testStore := StubUserStore{testUsers}
	testServer := api.NewSonyServer(&testStore)

	for _, user := range testUsers {
		t.Run(fmt.Sprintf("Test User %s", user.Name), func(t *testing.T) {
			response := httptest.NewRecorder()
			testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, fmt.Sprint(user.ID)))

			wantedResponse, gotResponse := getExpectedAndReceivedUserAchievementValidResponse(t, user, response, pkg.GoldAchievementLevel)

			assertHttpResponseStatus(t, response.Code, http.StatusOK)
			assertContentType(t, response, pkg.JsonContentType)
			assertResponseBody(t, gotResponse, wantedResponse)
		})
	}
}

func TestGetUserAchievementLevelForPlatinumAchievementLevelUsers(t *testing.T) {
	testUsers := []UserData{
		// Platinum Achievement Level Users:
		// Owns 50+ games and has 100% achievements in each

		// User owns 51 Games and has completed 100% achievements in all games.
		CustomNewTestUser(1, "Garry", map[string]int{
			"numberOfGames":         51,
			"numberOfAchievements":  100,
			"completedAchievements": 100,
		}),

		// User owns 100 Games and has completed 100% achievements in all games.
		CustomNewTestUser(2, "Tom", map[string]int{
			"numberOfGames":         100,
			"numberOfAchievements":  100,
			"completedAchievements": 100,
		}),
	}
	testStore := StubUserStore{testUsers}
	testServer := api.NewSonyServer(&testStore)

	for _, user := range testUsers {
		t.Run(fmt.Sprintf("Test User %s", user.Name), func(t *testing.T) {
			response := httptest.NewRecorder()
			testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, fmt.Sprint(user.ID)))

			wantedResponse, gotResponse := getExpectedAndReceivedUserAchievementValidResponse(t, user, response, pkg.PlatinumAchievementLevel)

			assertHttpResponseStatus(t, response.Code, http.StatusOK)
			assertContentType(t, response, pkg.JsonContentType)
			assertResponseBody(t, gotResponse, wantedResponse)
		})
	}
}

func TestAmbiguousBehaviour(t *testing.T) {
	testUsers := []UserData{
		CustomNewTestUser(1, "Garry", map[string]int{
			"numberOfGames":         15,
			"numberOfAchievements":  100,
			"completedAchievements": 100,
		}),
	}

	testStore := StubUserStore{testUsers}
	testServer := api.NewSonyServer(&testStore)

	t.Run("Check achievement level for incompatible user Id", func(t *testing.T) {
		response := httptest.NewRecorder()
		testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, "1a"))

		wantedErrorResponse, gotErrorResponse := getExpectedAndReceivedUserAchievementErrorResponse(t, response, pkg.ErrInvalidUserID)

		assertHttpResponseStatus(t, response.Code, http.StatusBadRequest)
		assertContentType(t, response, pkg.JsonContentType)
		assertResponseBody(t, gotErrorResponse, wantedErrorResponse)
	})

	t.Run("Check achievement level for non existing user", func(t *testing.T) {
		response := httptest.NewRecorder()
		testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, "2"))

		wantedErrorResponse, gotErrorResponse := getExpectedAndReceivedUserAchievementErrorResponse(t, response, pkg.ErrUserNotFound)

		assertHttpResponseStatus(t, response.Code, http.StatusNotFound)
		assertContentType(t, response, pkg.JsonContentType)
		assertResponseBody(t, gotErrorResponse, wantedErrorResponse)
	})
}

func TestCustomUser(t *testing.T) {
	testUsers := []UserData{
		CustomNewTestUser(1, "Garry",
			map[string]int{
				"numberOfGames":         1,
				"numberOfAchievements":  100,
				"completedAchievements": 100,
			}, map[string]int{
				"numberOfGames":         1,
				"numberOfAchievements":  100,
				"completedAchievements": 75,
			}, map[string]int{
				"numberOfGames":         1,
				"numberOfAchievements":  100,
				"completedAchievements": 50,
			}),
	}

	testStore := StubUserStore{testUsers}
	testServer := api.NewSonyServer(&testStore)

	for _, user := range testUsers {
		t.Run(fmt.Sprintf("Test User %s", user.Name), func(t *testing.T) {
			response := httptest.NewRecorder()
			testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, fmt.Sprint(user.ID)))

			wantedResponse, gotResponse := getExpectedAndReceivedUserAchievementValidResponse(t, user, response, pkg.NoAchievementLevel)

			assertHttpResponseStatus(t, response.Code, http.StatusOK)
			assertContentType(t, response, pkg.JsonContentType)
			assertResponseBody(t, gotResponse, wantedResponse)
		})
	}
}
