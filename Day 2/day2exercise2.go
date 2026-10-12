package main

import "fmt"

func calculateOrderValue(price float64, quantity int, discount float64) float64 {
	gross := price * float64(quantity)
	discountAmount := gross * (discount / 100)
	return gross - discountAmount
}

func main() {
	price := 150.0
	quantity := 4
	discount := 10.0 // 10% discount

	finalValue := calculateOrderValue(price, quantity, discount)
	fmt.Printf("Final Order Value: $%.2f\n", finalValue)
}
