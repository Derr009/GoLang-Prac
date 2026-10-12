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
		return errors.New("ID cannot be empty")
	}
	if t.Amount < 0 {
		return fmt.Errorf("negative amount: %.2f", t.Amount)
	}
	if t.Status != "SUCCESS" && t.Status != "FAILED" && t.Status != "CANCELLED" && t.Status != "PENDING" {
		return fmt.Errorf("invalid status: '%s'", t.Status)
	}
	return nil
}

func ProcessRecords(records []Transaction) ([]Transaction, []Transaction, []string) {
	var valid []Transaction
	var invalid []Transaction
	var errMsgs []string

	for index, t := range records {
		err := ValidateTransaction(t)
		if err != nil {
			invalid = append(invalid, t)
			errMsg := fmt.Sprintf("Record #%d [ID: %s]: %v", index+1, t.ID, err)
			errMsgs = append(errMsgs, errMsg)
		} else {
			valid = append(valid, t)
		}
	}

	return valid, invalid, errMsgs
}

func main() {
	records := []Transaction{
		{"T001", 100.0, "SUCCESS"},
		{"", 250.0, "SUCCESS"},
		{"T003", -40.0, "FAILED"},
		{"T004", 500.0, "CANCELLED"},
		{"T005", 120.0, "INVALID_STATUS"},
		{"T006", 0.0, "PENDING"},
		{"T007", 750.0, "SUCCESS"},
		{"T008", -10.0, "PENDING"},
		{"", -100.0, "BAD"},
		{"T010", 310.0, "SUCCESS"},
	}

	valid, invalid, errMsgs := ProcessRecords(records)

	fmt.Printf("Total Records Processed: %d\n", len(records))
	fmt.Printf("Valid Records: %d\n", len(valid))
	fmt.Printf("Invalid Records: %d\n", len(invalid))

	fmt.Println("\n--- Error Log ---")
	for _, msg := range errMsgs {
		fmt.Println(" -", msg)
	}
}
