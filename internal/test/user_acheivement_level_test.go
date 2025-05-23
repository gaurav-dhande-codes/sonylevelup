package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sonylevelup/internal/api"
)

func TestGetUserAcheivementLevel(t *testing.T) {
	t.Run("return garrys achievement level", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/user/1/acheivement-level", nil)
		response := httptest.NewRecorder()

		api.SonyServer(response, request)

		got := response.Body.String()
		want := "Bronze"

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})

	t.Run("return sallys achievement level", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/user/2/achievement-level", nil)
		response := httptest.NewRecorder()

		api.SonyServer(response, request)

		got := response.Body.String()
		want := "Silver"

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})
}
