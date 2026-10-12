package main

import "fmt"

// Define a simple struct for clean data organization
type Transaction struct {
	ID     string
	Amount float64
	Status string
}

func analyze(transactions []Transaction) (float64, float64, int, int, float64) {
	total := 0.0
	successCount := 0
	failedCount := 0
	highest := 0.0

	for _, t := range transactions {
		total += t.Amount

		if t.Amount > highest {
			highest = t.Amount
		}

		switch t.Status {
		case "SUCCESS":
			successCount++
		case "FAILED":
			failedCount++
		}
	}

	avg := 0.0
	if len(transactions) > 0 {
		avg = total / float64(len(transactions))
	}

	return total, avg, successCount, failedCount, highest
}

func main() {
	transactions := []Transaction{
		{"T001", 500.0, "SUCCESS"},
		{"T002", 120.0, "FAILED"},
		{"T003", 750.0, "SUCCESS"},
		{"T004", 1000.0, "SUCCESS"},
	}

	total, avg, successCount, failedCount, highest := analyze(transactions)

	fmt.Println("=== Transaction Analysis Summary ===")
	fmt.Printf("Total Amount: %.2f\n", total)
	fmt.Printf("Average Amount: %.2f\n", avg)
	fmt.Printf("Successful Count: %d\n", successCount)
	fmt.Printf("Failed Count: %d\n", failedCount)
	fmt.Printf("Highest Transaction: %.2f\n", highest)
}
