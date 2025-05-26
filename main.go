package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sonylevelup/internal/api"
)

func main() {
	fmt.Println("Welcome to Sony Level Up!")

	mockServerUserStore := api.NewMockServerUserStore("http://localhost:8080")
	server := api.NewSonyServer(mockServerUserStore)

	log.Fatal(http.ListenAndServe(":5000", server))
}
