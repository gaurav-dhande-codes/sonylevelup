package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sonylevelup/internal/api"
	"github.com/sonylevelup/internal/pkg"
)

func TestGetAllUsers(t *testing.T) {
	t.Run("Check valid response for no users in user store", func(t *testing.T) {
		testUsers := []UserData{}
		testStore := StubUserStore{testUsers}
		testServer := api.NewSonyServer(&testStore)
		response := httptest.NewRecorder()

		testServer.ServeHTTP(response, newGetAllUsesrRequest(t))

		wantedResponse, gotResponse := getExpectedAndReceivedGetAllUsersValidResponse(t, testUsers, response)

		assertHttpResponseStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, pkg.JsonContentType)
		assertResponseBody(t, gotResponse, wantedResponse)
	})

	t.Run("Check valid response for users in user store", func(t *testing.T) {
		testUsers := []UserData{
			CustomNewTestUser(1, "Garry", map[string]int{
				"numberOfGames":         0,
				"numberOfAchievements":  0,
				"completedAchievements": 0,
			}),
		}

		testStore := StubUserStore{testUsers}
		testServer := api.NewSonyServer(&testStore)
		response := httptest.NewRecorder()

		testServer.ServeHTTP(response, newGetAllUsesrRequest(t))

		wantedResponse, gotResponse := getExpectedAndReceivedGetAllUsersValidResponse(t, testUsers, response)

		assertHttpResponseStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, pkg.JsonContentType)
		assertResponseBody(t, gotResponse, wantedResponse)
	})

}
