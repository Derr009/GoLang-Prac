package main

import "fmt"

type Transaction struct {
	ID     string
	Amount float64
	Status string
}

func main() {
	transactions := []Transaction{
		{"T001", 150.0, "SUCCESS"},
		{"T002", 620.0, "SUCCESS"},
		{"T003", 200.0, "FAILED"},
		{"T004", 850.0, "SUCCESS"},
		{"T005", 100.0, "CANCELLED"},
		{"T006", 300.0, "FAILED"},
	}

	// Map key: Status, Map value: Slice of Transactions
	grouped := make(map[string][]Transaction)

	for _, t := range transactions {
		grouped[t.Status] = append(grouped[t.Status], t)
	}

	fmt.Println("=== Grouped Transactions ===")
	for status, txns := range grouped {
		fmt.Printf("\nStatus: %s (Count: %d)\n", status, len(txns))
		for _, t := range txns {
			fmt.Printf("  - ID: %s | Amount: %.2f\n", t.ID, t.Amount)
		}
	}
}
