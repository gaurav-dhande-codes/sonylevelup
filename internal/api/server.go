package api

import (
	"fmt"
	"net/http"
	"strings"
)

func SonyServer(w http.ResponseWriter, r *http.Request) {
	userId := strings.Split(r.URL.String(), "/")[2]

	if userId == "1" {
		fmt.Fprintf(w, "Bronze")
	}

	if userId == "2" {
		fmt.Fprintf(w, "Silver")
	}
}
