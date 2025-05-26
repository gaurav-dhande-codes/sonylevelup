package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sonylevelup/internal/model"
	"github.com/sonylevelup/internal/pkg"
)

// MockServerUserStore is a mock implementation of a user data store that fetches user-related data from a remote mock server via HTTP requests.
// It uses a base URL to construct API endpoints for retrieving user info, game library, and achievement completion data.
type MockServerUserStore struct {
	baseUrl string
}

// NewMockServerUserStore creates a new MockServerUserStore instance with the specified base URL for the mock server API.
func NewMockServerUserStore(baseUrl string) *MockServerUserStore {
	return &MockServerUserStore{baseUrl: baseUrl}
}

// GetUser retrieves a User object by userId by making an HTTP GET request to the mock server.
// It returns the user data on success or an error if the user is not found or any other failure occurs.
func (m *MockServerUserStore) GetUser(userId int) (*model.User, error) {
	getUserUrl := fmt.Sprintf("%s/users/%d", m.baseUrl, userId)

	response, err := http.Get(getUserUrl)
	if err != nil {
		return nil, fmt.Errorf("error while making get user request to mock server, %v", err)
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusNotFound:
		return nil, pkg.ErrUserNotFound
	case http.StatusOK:
		// continue as usual
	default:
		return nil, fmt.Errorf("error response received while making get user request to mock server, %v", response.StatusCode)
	}

	user := &model.User{}
	err = json.NewDecoder(response.Body).Decode(user)
	if err != nil {
		return nil, fmt.Errorf("error encountered while decoding get user response received from mock server, %v", err)
	}

	return user, nil
}

// GetUserGameLibrary fetches the game library for a given userId from the mock server via an HTTP GET request.
// It returns the user's game library on success or an error if the user is not found or the request fails.
func (m *MockServerUserStore) GetUserGameLibrary(userId int) (*model.UserLibrary, error) {
	getUserGameLibraryUrl := fmt.Sprintf("%s/users/%d/library", m.baseUrl, userId)

	response, err := http.Get(getUserGameLibraryUrl)
	if err != nil {
		return nil, fmt.Errorf("error while making get user library request to mock server, %v", err)
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusNotFound:
		return nil, pkg.ErrUserNotFound
	case http.StatusOK:
		// continue as usual
	default:
		return nil, fmt.Errorf("error response received while making get user library request to mock server, %v", response.StatusCode)
	}

	userLibrary := &model.UserLibrary{}
	err = json.NewDecoder(response.Body).Decode(userLibrary)
	if err != nil {
		return nil, fmt.Errorf("error encountered while decoding get user library response received from mock server, %v", err)
	}

	return userLibrary, nil
}

// GetUserGameAchievementCompletion retrieves the achievement completion details for a specific user and game from the mock server.
// It performs an HTTP GET request and returns the completion data or an error if the user or game is not found or if the request fails.
func (m *MockServerUserStore) GetUserGameAchievementCompletion(userId, gameId int) (*model.UserGameAchievementCompletion, error) {
	getUserGameAchievementCompletionUrl := fmt.Sprintf("%s/users/%d/achievements/%d", m.baseUrl, userId, gameId)

	response, err := http.Get(getUserGameAchievementCompletionUrl)
	if err != nil {
		return nil, fmt.Errorf("error while making get user game achievement completion request to mock server, %v", err)
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusNotFound:
		return nil, pkg.ErrUserOrGameNotFound
	case http.StatusOK:
		// continue as usual
	default:
		return nil, fmt.Errorf("error response received while making get user game achievement completion request to mock server, %v", response.StatusCode)
	}

	userGameAchievementCompletion := &model.UserGameAchievementCompletion{}
	err = json.NewDecoder(response.Body).Decode(userGameAchievementCompletion)
	if err != nil {
		return nil, fmt.Errorf("error encountered while decoding get user game achievement completion response received from mock server, %v", err)
	}

	return userGameAchievementCompletion, nil
}

// GetAllUser retrieves all users from the mock server.
// It performs an HTTP Get request and return a slice of users or an error if the request fails.
func (m *MockServerUserStore) GetAllUsers() ([]*model.User, error) {
	getAllUsersUrl := fmt.Sprintf("%s/users", m.baseUrl)

	response, err := http.Get(getAllUsersUrl)
	if err != nil {
		return nil, fmt.Errorf("error while making get all users request to mock server, %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response received while making get all users request to mock server, %v", response.StatusCode)
	}

	allUsers := []*model.User{}
	err = json.NewDecoder(response.Body).Decode(&allUsers)
	if err != nil {
		return nil, fmt.Errorf("error encountered while decoding get all users response received from mock server, %v", err)
	}

	return allUsers, nil
}
