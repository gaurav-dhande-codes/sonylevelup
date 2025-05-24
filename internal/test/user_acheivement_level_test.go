package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sonylevelup/internal/api"
)

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

func TestGetUserAchievementLevel(t *testing.T) {
	testUsers := []UserData{
		// No Rank Level Users:
		// 10 or fewer games owned
		NewTestUser(1, "Garry", 0, 0, 0),
		NewTestUser(2, "Tom", 1, 100, 100),
		NewTestUser(3, "Bob", 10, 100, 80),
		NewTestUser(4, "Luna", 10, 100, 75),
		NewTestUser(5, "Jerry", 10, 100, 75),
	}
	testStore := StubUserStore{testUsers}
	testServer := api.NewSonyServer(&testStore)

	t.Run("Test users who own 10 games or fewer", func(t *testing.T) {
		response := httptest.NewRecorder()
		testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, "1"))

		got := response.Body.String()
		want := "No Achievement Assigned"

		assertResponseBody(t, got, want)
	})

	t.Run("Test users who own 10 games or fewer", func(t *testing.T) {
		response := httptest.NewRecorder()
		testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, "2"))

		got := response.Body.String()
		want := "No Achievement Assigned"

		assertResponseBody(t, got, want)
	})
	t.Run("Test users who own 10 games or fewer", func(t *testing.T) {
		response := httptest.NewRecorder()
		testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, "3"))

		got := response.Body.String()
		want := "No Achievement Assigned"

		assertResponseBody(t, got, want)
	})
	t.Run("Test users who own 10 games or fewer", func(t *testing.T) {
		response := httptest.NewRecorder()
		testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, "4"))

		got := response.Body.String()
		want := "No Achievement Assigned"

		assertResponseBody(t, got, want)
	})

	t.Run("Test users who own 10 games or fewer", func(t *testing.T) {
		response := httptest.NewRecorder()
		testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, "5"))

		got := response.Body.String()
		want := "No Achievement Assigned"

		assertResponseBody(t, got, want)
	})
}

func newGetUserAchievementLevelRequest(t testing.TB, userId string) *http.Request {
	t.Helper()
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/user/%s/achievement-level", userId), nil)

	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
