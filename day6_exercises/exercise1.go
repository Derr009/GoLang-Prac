package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Call a public GET API (JSONPlaceholder)
	url := "https://jsonplaceholder.typicode.com/users/1"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("HTTP GET request failed: %v", err)
	}
	defer resp.Body.Close()

	// Read raw body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Print Status Code and Raw Response
	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Println("--- Response Body ---")
	fmt.Println(string(body))
}
