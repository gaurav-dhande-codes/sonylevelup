package api

import (
	"net/http"

	"github.com/sonylevelup/internal/pkg"
)

// GetAllUsers retrieves all users from the data store.
func (s *SonyServer) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Fetch all users data from the data store
	users, err := s.store.GetAllUsers()
	if err != nil {
		pkg.ErrorLogger.Printf("Failed unexpectedly when fetching all users from store: %v", err)
		pkg.WriteErrorResponseToResponseWriter(w, http.StatusInternalServerError, pkg.ErrInternalServerError)
		return
	}

	WriteValidResponseToResponseWriter(w, http.StatusOK, users)
}
