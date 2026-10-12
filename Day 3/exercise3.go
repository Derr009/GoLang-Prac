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
		{"T007", 300.0, "FAILED"},
		{"T008", 950.0, "SUCCESS"},
		{"T009", 400.0, "SUCCESS"},
		{"T010", 1200.0, "SUCCESS"},
	}

	fmt.Println("=== 10 Transactions Created ===")
	for _, t := range transactions {
		fmt.Printf("ID: %s | Amount: %.2f | Status: %s\n", t.ID, t.Amount, t.Status)
	}
}
