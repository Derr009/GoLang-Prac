package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Summary struct {
	Total      int `json:"total"`
	Valid      int `json:"valid"`
	Invalid    int `json:"invalid"`
	Duplicates int `json:"duplicates"`
}

func validateRecord(id, amountStr, status, dateStr string) error {
	if id == "" {
		return fmt.Errorf("missing ID")
	}

	amt, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amt <= 0 {
		return fmt.Errorf("invalid amount '%s'", amountStr)
	}

	if status != "SUCCESS" && status != "FAILED" && status != "CANCELLED" && status != "PENDING" {
		return fmt.Errorf("invalid status '%s'", status)
	}

	// Validate date format (YYYY-MM-DD)
	_, err = time.Parse("2006-01-02", dateStr)
	if err != nil {
		return fmt.Errorf("invalid date '%s' (expected YYYY-MM-DD)", dateStr)
	}

	return nil
}

func main() {
	inputFile, err := os.Open("Transactions.csv")
	if err != nil {
		log.Fatalf("Failed to open Transactions.csv: %v", err)
	}
	defer inputFile.Close()

	reader := csv.NewReader(inputFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV: %v", err)
	}

	var header []string
	if len(records) > 0 {
		header = records[0]
	}

	seenIDs := make(map[string]bool)
	var cleanRecords [][]string
	var invalidRecords [][]string

	duplicateCount := 0
	validCount := 0
	invalidCount := 0
	totalCount := len(records) - 1 // Exclude header

	if totalCount < 0 {
		totalCount = 0
	}

	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}

		id := record[0]
		amountStr := record[1]
		status := record[2]
		dateStr := record[3]

		// Check for duplicate ID
		if id != "" && seenIDs[id] {
			duplicateCount++
			invalidCount++
			invalidRow := append(record, "Duplicate ID")
			invalidRecords = append(invalidRecords, invalidRow)
			continue
		}

		// Validate row fields
		err := validateRecord(id, amountStr, status, dateStr)
		if err != nil {
			invalidCount++
			invalidRow := append(record, err.Error())
			invalidRecords = append(invalidRecords, invalidRow)
		} else {
			seenIDs[id] = true
			validCount++
			cleanRecords = append(cleanRecords, record)
		}
	}

	// 1. Write clean_transactions.csv
	cleanFile, err := os.Create("clean_transactions.csv")
	if err == nil {
		writer := csv.NewWriter(cleanFile)
		writer.Write(header)
		writer.WriteAll(cleanRecords)
		writer.Flush()
		cleanFile.Close()
	}

	// 2. Write invalid_transactions.csv
	invalidFile, err := os.Create("invalid_transactions.csv")
	if err == nil {
		writer := csv.NewWriter(invalidFile)
		invalidHeader := append(header, "Reason")
		writer.Write(invalidHeader)
		writer.WriteAll(invalidRecords)
		writer.Flush()
		invalidFile.Close()
	}

	// 3. Write Summary.json
	summary := Summary{
		Total:      totalCount,
		Valid:      validCount,
		Invalid:    invalidCount,
		Duplicates: duplicateCount,
	}

	summaryData, err := json.MarshalIndent(summary, "", "  ")
	if err == nil {
		os.WriteFile("Summary.json", summaryData, 0644)
	}

	fmt.Println("=== CSV Data Quality Checker Completed ===")
	fmt.Printf("Total: %d | Valid: %d | Invalid: %d | Duplicates: %d\n",
		summary.Total, summary.Valid, summary.Invalid, summary.Duplicates)
}
