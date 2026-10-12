package main

import (
	"fmt"
	"net/http"
)

func fetchEndpoint(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[%s] Request Error: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK: // 200
		fmt.Printf("[%s] Success (200 OK)\n", url)
	case http.StatusBadRequest: // 400
		fmt.Printf("[%s] Client Error (400 Bad Request)\n", url)
	case http.StatusNotFound: // 404
		fmt.Printf("[%s] Not Found (404)\n", url)
	case http.StatusInternalServerError: // 500
		fmt.Printf("[%s] Server Error (500 Internal Server Error)\n", url)
	default:
		fmt.Printf("[%s] Received Unexpected Status: %d\n", url, resp.StatusCode)
	}
}

func main() {
	// Testing various status codes using httpstat.us mock service
	urls := []string{
		"https://httpstat.us/200",
		"https://httpstat.us/400",
		"https://httpstat.us/404",
		"https://httpstat.us/500",
	}

	fmt.Println("=== Status Code Handler Test ===")
	for _, url := range urls {
		fetchEndpoint(url)
	}
}
