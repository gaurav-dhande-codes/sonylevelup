package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sonylevelup/internal/model"
	"github.com/sonylevelup/internal/pkg"
)

type MockServerUserStore struct {
	baseUrl string
}

func NewMockServerUserStore(baseUrl string) *MockServerUserStore {
	return &MockServerUserStore{baseUrl: baseUrl}
}

func (m *MockServerUserStore) GetUser(userId int) (*model.User, error) {
	getUserUrl := fmt.Sprintf("%s/users/%d", m.baseUrl, userId)
	fmt.Println(getUserUrl)

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
		return nil, fmt.Errorf("error response received while making get user request to mock server, %v", err)
	}

	user := &model.User{}
	err = json.NewDecoder(response.Body).Decode(user)
	if err != nil {
		return nil, fmt.Errorf("error encountered while decoding get user response received from mock server, %v", err)
	}

	return user, nil
}

func (m *MockServerUserStore) GetUserGameLibrary(userId int) (*model.UserLibrary, error) {
	return &model.UserLibrary{}, nil
}

func (m *MockServerUserStore) GetUserGameAchievementCompletion(userId, gameId int) (*model.UserGameAchievementCompletion, error) {
	return &model.UserGameAchievementCompletion{}, nil
}
