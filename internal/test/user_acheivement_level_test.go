package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sonylevelup/internal/api"
	"github.com/sonylevelup/internal/pkg"
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

func TestGetUserAchievementLevelForNoAchievementLevelUsers(t *testing.T) {
	testUsers := []UserData{
		// No Achievement Level Users:
		// 10 or fewer games owned
		NewTestUser(1, "Garry", 0, 0, 0),
		NewTestUser(2, "Tom", 1, 100, 100),
		NewTestUser(3, "Bob", 10, 100, 80),
		NewTestUser(4, "Luna", 10, 100, 75),
		NewTestUser(5, "Jerry", 10, 100, 75),
	}
	testStore := StubUserStore{testUsers}
	testServer := api.NewSonyServer(&testStore)

	for _, test := range testUsers {
		t.Run(fmt.Sprintf("Test User %s", test.Name), func(t *testing.T) {
			response := httptest.NewRecorder()
			testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, fmt.Sprint(test.ID)))

			got := response.Body.String()
			want := pkg.NoAchievementLevel

			assertResponseBody(t, got, want)
		})
	}
}

func TestGetUserAchievementLevelForBronzeAchievementLevelUsers(t *testing.T) {
	testUsers := []UserData{
		// No Rank Level Users:
		// More than 10 games owned
		// Atleast one game has achievement completion percentage of less than 75%
		NewTestUser(1, "Garry", 11, 100, 50),
		NewTestUser(2, "Tom", 15, 100, 60),
		NewTestUser(3, "Bob", 25, 100, 65),
		NewTestUser(4, "Luna", 30, 100, 70),
		NewTestUser(5, "Jerry", 75, 100, 75),
	}
	testStore := StubUserStore{testUsers}
	testServer := api.NewSonyServer(&testStore)

	for _, test := range testUsers {
		t.Run(fmt.Sprintf("Test User %s", test.Name), func(t *testing.T) {
			response := httptest.NewRecorder()
			testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, fmt.Sprint(test.ID)))

			got := response.Body.String()
			want := pkg.BronzeAchievementLevel

			assertResponseBody(t, got, want)
		})
	}
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
