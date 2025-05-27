package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/sonylevelup/internal/model"
	"github.com/sonylevelup/internal/pkg"
)

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
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s/achievement-level", userId), nil)
	if err != nil {
		t.Fatalf("GetUserAchievementLevel request failed with error: %v", err)
	}

	return request
}

// newGetAllUsesrRequest creates a new HTTP GET request for retrieving all users.
//
// Parameters:
//   - t: The testing interface (typically *testing.T).
//
// Returns:
//   - A pointer to an http.Request targeting the /users endpoint.
//
// Example:
//
//	req := newGetAllUsesrRequest(t)
func newGetAllUsesrRequest(t testing.TB) *http.Request {
	request, err := http.NewRequest(http.MethodGet, "/users", nil)
	if err != nil {
		t.Fatalf("GetUserAchievementLevel request failed with error: %v", err)
	}

	return request
}

// assertResponseBody compares the actual and expected user's achievement level responses.
//
// Parameters:
//   - t: The testing interface.
//   - got: The actual response received.
//   - want: The expected response.
//
// This function fails the test if `got` and `want` objects are not equal.
func assertResponseBody(t testing.TB, got, want any) {
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

// getExpectedAndReceivedGetAllUsersValidResponse constructs the expected response and decodes the actual response
// from the response recorder for the GetAllUsers handler.
//
// Parameters:
//   - t: The testing interface.
//   - users: A slice of UserData representing the expected users returned by the API.
//   - response: The HTTP response recorder containing the actual response.
//
// Returns:
//   - A pointer to a slice of model.User representing the expected response.
//   - A pointer to a slice of model.User decoded from the actual HTTP response.
//
// Example:
//
//	want, got := getExpectedAndReceivedGetAllUsersValidResponse(t, testUsers, recorder)
func getExpectedAndReceivedGetAllUsersValidResponse(
	t testing.TB,
	users []UserData,
	response *httptest.ResponseRecorder) (*[]model.User, *[]model.User) {
	t.Helper()

	wantedResponse := []model.User{}
	for _, user := range users {
		wantedResponse = append(wantedResponse,
			model.User{
				Id:    user.ID,
				Name:  user.Name,
				Email: user.Email})
	}

	gotResponse := &[]model.User{}
	err := json.NewDecoder(response.Body).Decode(gotResponse)
	if err != nil {
		t.Fatalf("error while decoding response: %v", err)
	}

	return &wantedResponse, gotResponse
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

// assertContentType checks whether the Content-Type header in the HTTP response matches the expected value.
//
// Parameters:
//   - t: The testing interface (e.g., *testing.T or *testing.B).
//   - response: The HTTP response recorder used to capture the test response.
//   - want: The expected Content-Type header value (e.g., "application/json").
//
// This function fails the test if the Content-Type header does not match the expected value.
//
// Example:
//
//	assertContentType(t, recorder, "application/json")
func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response header did not have content-type of %s, got: %v", want, response.Result().Header)
	}
}
