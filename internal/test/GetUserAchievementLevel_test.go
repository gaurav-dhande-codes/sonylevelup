package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/sonylevelup/internal/api"
	"github.com/sonylevelup/internal/model"
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
			assertAchievementLevelResponse(t, gotResponse, wantedResponse)
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
			assertAchievementLevelResponse(t, gotResponse, wantedResponse)
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
			assertAchievementLevelResponse(t, gotResponse, wantedResponse)
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
			assertAchievementLevelResponse(t, gotResponse, wantedResponse)
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
			assertAchievementLevelResponse(t, gotResponse, wantedResponse)
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
		assertAchievementLevelResponse(t, gotErrorResponse, wantedErrorResponse)
	})

	t.Run("Check achievement level for non existing user", func(t *testing.T) {
		response := httptest.NewRecorder()
		testServer.ServeHTTP(response, newGetUserAchievementLevelRequest(t, "2"))

		wantedErrorResponse, gotErrorResponse := getExpectedAndReceivedUserAchievementErrorResponse(t, response, pkg.ErrUserNotFound)

		assertHttpResponseStatus(t, response.Code, http.StatusNotFound)
		assertAchievementLevelResponse(t, gotErrorResponse, wantedErrorResponse)
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
			assertAchievementLevelResponse(t, gotResponse, wantedResponse)
		})
	}
}

// newGetUserAchievementLevelRequest creates a new HTTP GET request for retrieving a user's achievement level.
//
// Parameters:
//   - t: The testing interface (typically *testing.T).
//   - userId: The user ID to include in the request path.
//
// Returns:
//   - A pointer to an http.Request targeting the achievement-level endpoint.
//
// Example:
//
//	req := newGetUserAchievementLevelRequest(t, "123")
func newGetUserAchievementLevelRequest(t testing.TB, userId string) *http.Request {
	t.Helper()
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s/achievement-level", userId), nil)

	return request
}

// assertAchievementLevelResponse compares the actual and expected user's achievement level responses.
//
// Parameters:
//   - t: The testing interface.
//   - got: The actual response received.
//   - want: The expected response.
//
// This function fails the test if `got` and `want` objects are not equal.
func assertAchievementLevelResponse(t testing.TB, got, want any) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got achievement level response body: %v, want: %v", got, want)
	}
}

// getExpectedAndReceivedUserAchievementErrorResponse decodes the actual error response from the HTTP response
// and constructs the expected one using a given error then compares them.
//
// Parameters:
//   - t: The testing interface.
//   - response: The HTTP response recorder containing the actual response.
//   - wantedError: The expected error used to create the expected error response.
//
// Returns:
//   - A pointer to the expected pkg.ErrorResponse.
//   - A pointer to the actual pkg.ErrorResponse decoded from the response body.
//
// Example:
//
//	want, got := getExpectedAndReceivedUserAchievementErrorResponse(t, resp, pkg.ErrUserNotFound)
func getExpectedAndReceivedUserAchievementErrorResponse(
	t testing.TB,
	response *httptest.ResponseRecorder,
	wantedError error) (*pkg.ErrorResponse, *pkg.ErrorResponse) {
	t.Helper()

	gotErrorResponse := &pkg.ErrorResponse{}
	err := json.NewDecoder(response.Body).Decode(gotErrorResponse)
	if err != nil {
		t.Fatalf("error while decoding response: %v", err)
	}

	wantedErrorResponse := &pkg.ErrorResponse{
		Error:   pkg.ErrorMapping[wantedError],
		Message: wantedError.Error(),
	}

	return wantedErrorResponse, gotErrorResponse
}

// getExpectedAndReceivedUserAchievementValidResponse constructs the expected valid response for a user achievement level,
// and decodes the actual response from the response recorder.
//
// Parameters:
//   - t: The testing interface.
//   - user: The user whose achievement level is being tested.
//   - response: The HTTP response recorder containing the actual response.
//   - achievementLevel: The expected achievement level string.
//
// Returns:
//   - A pointer to the expected model.UserAchievementLevel.
//   - A pointer to the actual model.UserAchievementLevel decoded from the response.
//
// Example:
//
//	want, got := getExpectedAndReceivedUserAchievementValidResponse(t, user, resp, "Gold")
func getExpectedAndReceivedUserAchievementValidResponse(
	t testing.TB,
	user UserData,
	response *httptest.ResponseRecorder,
	achievementLevel string) (*model.UserAchievementLevel, *model.UserAchievementLevel) {
	t.Helper()

	wantedResponse := &model.UserAchievementLevel{
		User: &model.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email},
		AchievementLevel: achievementLevel,
	}

	gotResponse := &model.UserAchievementLevel{}
	err := json.NewDecoder(response.Body).Decode(gotResponse)
	if err != nil {
		t.Fatalf("error while decoding response: %v", err)
	}

	return wantedResponse, gotResponse
}

// assertHttpResponseStatus compares the received HTTP status code with the expected value.
//
// Parameters:
//   - t: The testing interface.
//   - got: The actual HTTP status code returned.
//   - want: The expected HTTP status code.
//
// This function fails the test if the status codes do not match.
//
// Example:
//
//	assertHttpResponseStatus(t, recorder.Code, http.StatusOK)
func assertHttpResponseStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("recieved incorrect status code, got: %d, want: %d", got, want)
	}
}
