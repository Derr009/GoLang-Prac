package main

import "fmt"

func computeBreakdown(price float64, quantity int, discountPct, taxPct float64) (float64, float64, float64, float64) {
	gross := price * float64(quantity)
	discountAmt := gross * (discountPct / 100.0)
	discountedGross := gross - discountAmt
	taxAmt := discountedGross * (taxPct / 100.0)
	finalAmt := discountedGross + taxAmt

	return gross, discountAmt, taxAmt, finalAmt
}

func main() {
	var product string
	var price float64
	var quantity int
	var discountPct, taxPct float64

	// User Inputs
	fmt.Print("Enter Product Name: ")
	fmt.Scanln(&product)

	fmt.Print("Enter Price: ")
	fmt.Scanln(&price)

	fmt.Print("Enter Quantity: ")
	fmt.Scanln(&quantity)

	fmt.Print("Enter Discount %: ")
	fmt.Scanln(&discountPct)

	fmt.Print("Enter Tax %: ")
	fmt.Scanln(&taxPct)

	// Compute values
	gross, discount, tax, final := computeBreakdown(price, quantity, discountPct, taxPct)

	// Display Output
	fmt.Println("\n=== Sales Summary ===")
	fmt.Printf("Product: %s\n", product)
	fmt.Printf("Gross Amount: %.2f\n", gross)
	fmt.Printf("Discount: %.2f\n", discount)
	fmt.Printf("Tax: %.2f\n", tax)
	fmt.Printf("Final Amount: %.2f\n", final)
}
