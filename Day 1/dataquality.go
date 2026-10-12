package main

import "fmt"

// Function to calculate rate
func calculateRate(count, total int) float64 {
	if total == 0 {
		return 0.0
	}
	return (float64(count) / float64(total)) * 100
}

func main() {
	var totalRecords, validRecords, invalidRecords, duplicateRecords int

	// User Input
	fmt.Print("Enter Total records: ")
	fmt.Scanln(&totalRecords)

	fmt.Print("Enter Valid records: ")
	fmt.Scanln(&validRecords)

	fmt.Print("Enter Invalid records: ")
	fmt.Scanln(&invalidRecords)

	fmt.Print("Enter Duplicate records: ")
	fmt.Scanln(&duplicateRecords)

	// Validation check
	if validRecords+invalidRecords+duplicateRecords > totalRecords {
		fmt.Println("Warning: Sum of valid, invalid, and duplicate records exceeds total records!")
	}

	// Rate Calculations
	validationRate := calculateRate(validRecords, totalRecords)
	errorRate := calculateRate(invalidRecords, totalRecords)
	duplicateRate := calculateRate(duplicateRecords, totalRecords)

	// Display Summary Output
	fmt.Println("\nData Quality Summary")
	fmt.Println("--------------------")
	fmt.Printf("Total Records: %d\n", totalRecords)
	fmt.Printf("Valid Records: %d\n", validRecords)
	fmt.Printf("Invalid Records: %d\n", invalidRecords)
	fmt.Printf("Duplicate Records: %d\n\n", duplicateRecords)

	fmt.Printf("Validation Rate: %.0f%%\n", validationRate)
	fmt.Printf("Error Rate: %.0f%%\n", errorRate)
	fmt.Printf("Duplicate Rate: %.0f%%\n", duplicateRate)
}
