package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Transaction struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

func main() {
	// 1. Read
	file, err := os.Open("input_transactions.csv")
	if err != nil {
		log.Fatalf("Step 1 Failed (Read): %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Step 2 Failed (Parse CSV): %v", err)
	}

	var validTxns []Transaction

	// 2 & 3. Parse & Validate
	for i, record := range records {
		if i == 0 {
			continue // Skip Header
		}

		id := record[0]
		amountStr := record[1]
		status := record[2]

		if id == "" {
			fmt.Printf("Row %d Skipped: Empty ID\n", i)
			continue
		}

		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil || amount <= 0 {
			fmt.Printf("Row %d Skipped: Invalid Amount '%s'\n", i, amountStr)
			continue
		}

		// 4. Transform
		normalizedStatus := strings.ToUpper(strings.TrimSpace(status))

		validTxns = append(validTxns, Transaction{
			ID:     id,
			Amount: amount,
			Status: normalizedStatus,
		})
	}

	// 5. Write CSV
	csvFile, err := os.Create("processed_transactions.csv")
	if err != nil {
		log.Fatalf("Step 5 Failed (Write CSV): %v", err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	writer.Write([]string{"ID", "Amount", "Status"})

	for _, t := range validTxns {
		writer.Write([]string{t.ID, fmt.Sprintf("%.2f", t.Amount), t.Status})
	}
	writer.Flush()

	// 6. Write JSON
	jsonData, err := json.MarshalIndent(validTxns, "", "  ")
	if err != nil {
		log.Fatalf("Step 6 Failed (Marshal JSON): %v", err)
	}

	err = os.WriteFile("processed_transactions.json", jsonData, 0644)
	if err != nil {
		log.Fatalf("Step 6 Failed (Write JSON): %v", err)
	}

	fmt.Println("=== Guided Pipeline Successfully Executed ===")
	fmt.Printf("Processed Records: %d\n", len(validTxns))
}
