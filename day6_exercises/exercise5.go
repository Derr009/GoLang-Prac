package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetchWithTimeout(url string, timeout time.Duration) {
	// Create custom HTTP client with timeout settings
	client := &http.Client{
		Timeout: timeout,
	}

	fmt.Printf("Fetching %s with %v timeout...\n", url, timeout)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error: Request failed or timed out: %v\n\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Success! Status Code: %d\n\n", resp.StatusCode)
}

func main() {
	// Fast request that will succeed within 2 seconds
	fetchWithTimeout("https://jsonplaceholder.typicode.com/users/1", 2*time.Second)

	// Simulated slow endpoint (delays response by 3 seconds) with a 1 second timeout to trigger timeout error
	fetchWithTimeout("https://httpstat.us/200?sleep=3000", 1*time.Second)
}
