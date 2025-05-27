package main

import (
	"fmt"
	"net/http"

	"github.com/sonylevelup/internal/api"
	"github.com/sonylevelup/internal/pkg"
)

func main() {
	fmt.Println("Welcome to Sony Level Up!")

	mockServerUserStore := api.NewMockServerUserStore("http://localhost:8080")
	server := api.NewSonyServer(mockServerUserStore)

	pkg.ErrorLogger.Fatal(http.ListenAndServe(":5000", server))
}
