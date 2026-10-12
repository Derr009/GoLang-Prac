package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Raw API response struct (JSON Unmarshaling target)
type RawUser struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Company  struct {
		Name string `json:"name"`
	} `json:"company"`
}

// Clean Transformed Struct for CSV export
type CleanUser struct {
	ID          int
	Name        string
	Email       string
	CompanyName string
	Status      string
}

// Step 1 & 2: HTTP Request with Timeout & Status Code Validation
func fetchUsers(url string, timeout time.Duration) ([]RawUser, error) {
	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Step 3: Parse JSON into Struct
	var users []RawUser
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return users, nil
}

// Step 4 & 5: Validate & Transform
func processUsers(rawUsers []RawUser) []CleanUser {
	var cleanUsers []CleanUser

	for _, raw := range rawUsers {
		// Validation logic: Skip record if mandatory fields are missing
		if raw.ID <= 0 || strings.TrimSpace(raw.Name) == "" || !strings.Contains(raw.Email, "@") {
			fmt.Printf("Skipping invalid user record: %+v\n", raw)
			continue
		}

		// Transformation logic: Upper case company name, sanitize fields
		clean := CleanUser{
			ID:          raw.ID,
			Name:        strings.TrimSpace(raw.Name),
			Email:       strings.ToLower(strings.TrimSpace(raw.Email)),
			CompanyName: strings.ToUpper(strings.TrimSpace(raw.Company.Name)),
			Status:      "ACTIVE",
		}

		cleanUsers = append(cleanUsers, clean)
	}

	return cleanUsers
}

// Step 6: Export Clean Data to CSV
func exportToCSV(filename string, users []CleanUser) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Header row
	header := []string{"ID", "Name", "Email", "Company Name", "Status"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Data rows
	for _, u := range users {
		row := []string{
			fmt.Sprintf("%d", u.ID),
			u.Name,
			u.Email,
			u.CompanyName,
			u.Status,
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	apiURL := "https://jsonplaceholder.typicode.com/users"
	outputFile := "extracted_users.csv"

	fmt.Println("=== Starting API Data Extractor Pipeline ===")

	// 1. Fetch & Parse JSON
	fmt.Println("[1/4] Fetching data from API...")
	rawUsers, err := fetchUsers(apiURL, 5*time.Second)
	if err != nil {
		log.Fatalf("Error during fetch: %v", err)
	}
	fmt.Printf(" -> Fetched %d raw records\n", len(rawUsers))

	// 2. Validate & Transform
	fmt.Println("[2/4] Validating and transforming data...")
	cleanUsers := processUsers(rawUsers)
	fmt.Printf(" -> Successfully processed %d valid records\n", len(cleanUsers))

	// 3. Write to CSV
	fmt.Println("[3/4] Exporting clean data to CSV...")
	err = exportToCSV(outputFile, cleanUsers)
	if err != nil {
		log.Fatalf("Error writing CSV: %v", err)
	}

	fmt.Printf("[4/4] Pipeline Complete! Output written to %s\n", outputFile)
}
