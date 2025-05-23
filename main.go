package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sonylevelup/internal/api"
)

func main() {
	fmt.Println("Welcome to Sony Level Up!")
	server := &api.SonyServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}
