package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Customer struct matching the Customer API endpoint (/users)
type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Transaction struct matching the Transaction API endpoint (/todos)
type RawTransaction struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Joined Record for CSV Export
type CustomerTransaction struct {
	CustomerID    int
	CustomerName  string
	CustomerEmail string
	TransactionID int
	Title         string
	Status        string
}

// Helper HTTP Fetcher with Timeout and Error Handling
func fetchAPI(url string, client *http.Client) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request error for %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API %s returned status code: %d", url, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

func main() {
	// Reusable HTTP Client with Timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	customerAPI := "https://jsonplaceholder.typicode.com/users"
	transactionAPI := "https://jsonplaceholder.typicode.com/todos"
	outputCSV := "Customer_transactions.csv"

	fmt.Println("=== Multi-API Data Collector Started ===")

	// 1. Fetch Customer Data
	fmt.Println("[1/4] Fetching Customer API...")
	custBody, err := fetchAPI(customerAPI, client)
	if err != nil {
		log.Fatalf("Error fetching customers: %v", err)
	}

	var customers []Customer
	if err := json.Unmarshal(custBody, &customers); err != nil {
		log.Fatalf("Error parsing Customer JSON: %v", err)
	}
	fmt.Printf(" -> Loaded %d customers\n", len(customers))

	// 2. Fetch Transaction Data
	fmt.Println("[2/4] Fetching Transaction API...")
	txnBody, err := fetchAPI(transactionAPI, client)
	if err != nil {
		log.Fatalf("Error fetching transactions: %v", err)
	}

	var rawTxns []RawTransaction
	if err := json.Unmarshal(txnBody, &rawTxns); err != nil {
		log.Fatalf("Error parsing Transaction JSON: %v", err)
	}
	fmt.Printf(" -> Loaded %d transactions\n", len(rawTxns))

	// 3. Build In-Memory Customer Lookup Map
	customerMap := make(map[int]Customer)
	for _, c := range customers {
		customerMap[c.ID] = c
	}

	// 4. Transform and Join Data
	fmt.Println("[3/4] Joining customer and transaction datasets...")
	var joinedData []CustomerTransaction

	for _, txn := range rawTxns {
		// Lookup customer details by UserID
		cust, exists := customerMap[txn.UserID]
		if !exists {
			fmt.Printf("Warning: Transaction %d has unknown customer ID %d. Skipping.\n", txn.ID, txn.UserID)
			continue
		}

		status := "PENDING"
		if txn.Completed {
			status = "SUCCESS"
		}

		joinedData = append(joinedData, CustomerTransaction{
			CustomerID:    cust.ID,
			CustomerName:  strings.TrimSpace(cust.Name),
			CustomerEmail: strings.ToLower(strings.TrimSpace(cust.Email)),
			TransactionID: txn.ID,
			Title:         strings.TrimSpace(txn.Title),
			Status:        status,
		})
	}

	// 5. Write Combined Results to CSV
	fmt.Println("[4/4] Exporting combined data to CSV...")
	file, err := os.Create(outputCSV)
	if err != nil {
		log.Fatalf("Failed to create CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write Header
	header := []string{"Customer ID", "Customer Name", "Customer Email", "Transaction ID", "Title / Description", "Status"}
	if err := writer.Write(header); err != nil {
		log.Fatalf("Error writing CSV header: %v", err)
	}

	// Write Rows
	for _, record := range joinedData {
		row := []string{
			strconv.Itoa(record.CustomerID),
			record.CustomerName,
			record.CustomerEmail,
			strconv.Itoa(record.TransactionID),
			record.Title,
			record.Status,
		}
		if err := writer.Write(row); err != nil {
			log.Fatalf("Error writing CSV row: %v", err)
		}
	}

	fmt.Printf("=== Success! Processed %d joined records into %s ===\n", len(joinedData), outputCSV)
}
