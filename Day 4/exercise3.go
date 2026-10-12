package main

import (
	"errors"
	"fmt"
)

type Transaction struct {
	ID     string
	Amount float64
	Status string
}

func ValidateTransaction(t Transaction) error {
	if t.ID == "" {
		return errors.New("validation failed: transaction ID cannot be empty")
	}
	if t.Amount < 0 {
		return fmt.Errorf("validation failed for ID %s: negative amount %.2f", t.ID, t.Amount)
	}
	if t.Status != "SUCCESS" && t.Status != "FAILED" && t.Status != "PENDING" && t.Status != "CANCELLED" {
		return fmt.Errorf("validation failed for ID %s: invalid status '%s'", t.ID, t.Status)
	}
	return nil
}

func main() {
	testTxns := []Transaction{
		{ID: "T001", Amount: 150.0, Status: "SUCCESS"},
		{ID: "", Amount: 200.0, Status: "SUCCESS"},
		{ID: "T003", Amount: -50.0, Status: "SUCCESS"},
		{ID: "T004", Amount: 100.0, Status: "UNKNOWN"},
	}

	for _, t := range testTxns {
		err := ValidateTransaction(t)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Valid Transaction: %+v\n", t)
		}
	}
}
