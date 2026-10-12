package main

import "fmt"

type Transaction struct {
	Amount float64
}

// IsHighValue returns true if the transaction amount exceeds 500
func (t Transaction) IsHighValue() bool {
	return t.Amount > 500
}

func main() {
	t1 := Transaction{Amount: 250.0}
	t2 := Transaction{Amount: 750.0}

	fmt.Printf("Txn 1 ($%.2f) High Value: %t\n", t1.Amount, t1.IsHighValue())
	fmt.Printf("Txn 2 ($%.2f) High Value: %t\n", t2.Amount, t2.IsHighValue())
}
