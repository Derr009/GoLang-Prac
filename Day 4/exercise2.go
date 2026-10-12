package main

import "fmt"

type Transaction struct {
	ID     string
	Amount float64
	Status string
}

// Pointer receiver allows modifying the struct instance directly
func (t *Transaction) UpdateStatus(newStatus string) {
	t.Status = newStatus
}

func main() {
	txn := Transaction{ID: "T001", Amount: 300.0, Status: "PENDING"}

	fmt.Printf("Before Update: %+v\n", txn)
	txn.UpdateStatus("SUCCESS")
	fmt.Printf("After Update:  %+v\n", txn)
}
