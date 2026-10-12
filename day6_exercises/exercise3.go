package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type ExtractedTodo struct {
	ID     int
	Name   string
	Status string
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var rawTodo Todo
	json.Unmarshal(body, &rawTodo)

	// Extract and map specific fields
	statusStr := "PENDING"
	if rawTodo.Completed {
		statusStr = "COMPLETED"
	}

	extracted := ExtractedTodo{
		ID:     rawTodo.ID,
		Name:   rawTodo.Title,
		Status: statusStr,
	}

	fmt.Println("=== Extracted Data ===")
	fmt.Printf("ID: %d\n", extracted.ID)
	fmt.Printf("Name: %s\n", extracted.Name)
	fmt.Printf("Status: %s\n", extracted.Status)
}
