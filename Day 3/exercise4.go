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
		{"T006", 550.0, "SUCCESS"},
	}

	var highValueTxns []Transaction

	for _, t := range transactions {
		if t.Amount > 500 {
			highValueTxns = append(highValueTxns, t)
		}
	}

	fmt.Println("=== Transactions with Amount > 500 ===")
	for _, t := range highValueTxns {
		fmt.Printf("ID: %s | Amount: %.2f | Status: %s\n", t.ID, t.Amount, t.Status)
	}
}
