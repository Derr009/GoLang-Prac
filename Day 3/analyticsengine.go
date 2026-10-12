package main

import "fmt"

type Transaction struct {
	ID     string
	Amount float64
	Status string
}

func analyzeTransactions(transactions []Transaction, threshold float64) {
	if len(transactions) == 0 {
		fmt.Println("No transactions available.")
		return
	}

	count := len(transactions)
	totalAmount := 0.0
	maxAmt := transactions[0].Amount
	minAmt := transactions[0].Amount

	statusCount := make(map[string]int)
	statusAmount := make(map[string]float64)
	var aboveThreshold []Transaction

	for _, t := range transactions {
		totalAmount += t.Amount

		if t.Amount > maxAmt {
			maxAmt = t.Amount
		}
		if t.Amount < minAmt {
			minAmt = t.Amount
		}

		statusCount[t.Status]++
		statusAmount[t.Status] += t.Amount

		if t.Amount > threshold {
			aboveThreshold = append(aboveThreshold, t)
		}
	}

	avgAmount := totalAmount / float64(count)

	fmt.Println("=== Transaction Analytics Engine Output ===")
	fmt.Printf("Total Count: %d\n", count)
	fmt.Printf("Total Amount: %.2f\n", totalAmount)
	fmt.Printf("Average Amount: %.2f\n", avgAmount)
	fmt.Printf("Maximum Amount: %.2f\n", maxAmt)
	fmt.Printf("Minimum Amount: %.2f\n", minAmt)

	fmt.Println("\n--- Status-wise Count ---")
	for status, c := range statusCount {
		fmt.Printf("  %s: %d\n", status, c)
	}

	fmt.Println("\n--- Status-wise Amount ---")
	for status, amt := range statusAmount {
		fmt.Printf("  %s: %.2f\n", status, amt)
	}

	fmt.Printf("\n--- Transactions Above Threshold (%.2f) ---\n", threshold)
	for _, t := range aboveThreshold {
		fmt.Printf("  ID: %s | Amount: %.2f | Status: %s\n", t.ID, t.Amount, t.Status)
	}
}

func main() {
	transactions := []Transaction{
		{"T001", 100.0, "SUCCESS"},
		{"T002", 250.0, "FAILED"},
		{"T003", 600.0, "SUCCESS"},
		{"T004", 1200.0, "SUCCESS"},
		{"T005", 50.0, "CANCELLED"},
		{"T006", 800.0, "FAILED"},
	}

	analyzeTransactions(transactions, 500.0)
}
