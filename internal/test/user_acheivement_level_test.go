package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sonylevelup/internal/api"
)

type StubUserStore struct {
	users                          []api.User
	userLibraries                  []api.UserLibrary
	userGameAchievementCompletions []api.UserGameAchievementCompletion
}

func (s *StubUserStore) GetUser(userId int) *api.User {
	return &api.User{}
}

func (s *StubUserStore) GetUserGameLibrary(userId int) *api.UserLibrary {
	return &api.UserLibrary{}
}

func (s *StubUserStore) GetUserGameAchievementCompletion(userId, gameId int) *api.UserGameAchievementCompletion {
	return &api.UserGameAchievementCompletion{}
}

func TestGetUserAchievementLevel(t *testing.T) {
	testStore := StubUserStore{
		users:                          []api.User{},
		userLibraries:                  []api.UserLibrary{},
		userGameAchievementCompletions: []api.UserGameAchievementCompletion{},
	}
	testServer := api.NewSonyServer(&testStore)

	t.Run("return garrys achievement level", func(t *testing.T) {
		response := httptest.NewRecorder()
		testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, "1"))

		got := response.Body.String()
		want := "Bronze"

		assertResponseBody(t, got, want)
	})

	t.Run("return sallys achievement level", func(t *testing.T) {
		response := httptest.NewRecorder()
		testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, "2"))

		got := response.Body.String()
		want := "Silver"

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
