package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// User struct matching the JSON structure from the API
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/users/1"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("HTTP GET request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read body: %v", err)
	}

	// Unmarshal JSON into struct
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatalf("JSON Unmarshal error: %v", err)
	}

	fmt.Printf("Parsed Struct: %+v\n", user)
}
