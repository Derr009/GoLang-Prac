package main

import (
	"fmt"
	"sort"
)

type Transaction struct {
	Customer      string
	TransactionID string
	Amount        float64
	Status        string
}

type CustomerMetrics struct {
	Customer          string
	TotalTransactions int
	TotalAmount       float64
	SuccessfulCount   int
	FailedCount       int
}

func main() {
	transactions := []Transaction{
		{"Customer A", "T001", 1000.0, "SUCCESS"},
		{"Customer A", "T002", 2000.0, "SUCCESS"},
		{"Customer A", "T003", 500.0, "FAILED"},
		{"Customer A", "T004", 5000.0, "SUCCESS"},
		{"Customer B", "T005", 1500.0, "SUCCESS"},
		{"Customer B", "T006", 3000.0, "SUCCESS"},
		{"Customer C", "T007", 7000.0, "SUCCESS"},
		{"Customer C", "T008", 2000.0, "FAILED"},
		{"Customer D", "T009", 12000.0, "SUCCESS"},
		{"Customer E", "T010", 4000.0, "SUCCESS"},
	}

	// Map to accumulate metrics per customer
	customerMap := make(map[string]*CustomerMetrics)

	for _, t := range transactions {
		if _, exists := customerMap[t.Customer]; !exists {
			customerMap[t.Customer] = &CustomerMetrics{Customer: t.Customer}
		}

		metrics := customerMap[t.Customer]
		metrics.TotalTransactions++
		metrics.TotalAmount += t.Amount

		if t.Status == "SUCCESS" {
			metrics.SuccessfulCount++
		} else if t.Status == "FAILED" {
			metrics.FailedCount++
		}
	}

	// 1. Display Summary per Customer
	fmt.Println("=== Customer Transaction Summaries ===")
	for _, metrics := range customerMap {
		fmt.Printf("\n%s\n", metrics.Customer)
		fmt.Printf("Total Transactions: %d\n", metrics.TotalTransactions)
		fmt.Printf("Total Amount: %.0f\n", metrics.TotalAmount)
		fmt.Printf("Successful: %d\n", metrics.SuccessfulCount)
		fmt.Printf("Failed: %d\n", metrics.FailedCount)
	}

	// 2. Extension: Top 3 Customers by Transaction Value
	var summaryList []CustomerMetrics
	for _, metrics := range customerMap {
		summaryList = append(summaryList, *metrics)
	}

	// Sort descending by TotalAmount
	sort.Slice(summaryList, func(i, j int) bool {
		return summaryList[i].TotalAmount > summaryList[j].TotalAmount
	})

	fmt.Println("\n========================================")
	fmt.Println("Top 3 Customers by Transaction Value:")
	fmt.Println("========================================")
	limit := 3
	if len(summaryList) < limit {
		limit = len(summaryList)
	}

	for i := 0; i < limit; i++ {
		fmt.Printf("%d. %s — Total Value: %.2f\n", i+1, summaryList[i].Customer, summaryList[i].TotalAmount)
	}
}
