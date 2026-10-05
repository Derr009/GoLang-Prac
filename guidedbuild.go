package main

import "fmt"

// Function to calculate percentages
func calculatePercentage(part, total float64) float64 {
	if total == 0 {
		return 0.0
	}
	return (part / total) * 100
}

func main() {
	// Constants
	const header = "=== Transaction Summary ==="

	// Variables
	var total, successful, failed int

	// User Input
	fmt.Print("Enter Total Transactions: ")
	fmt.Scanln(&total)

	fmt.Print("Enter Successful Transactions: ")
	fmt.Scanln(&successful)

	fmt.Print("Enter Failed Transactions: ")
	fmt.Scanln(&failed)

	// Simple validation using if
	if successful+failed > total {
		fmt.Println("Error: Successful + Failed cannot exceed Total Transactions!")
		return
	}

	// Type conversion from int to float64 for division
	totalF := float64(total)
	successfulF := float64(successful)
	failedF := float64(failed)

	// Calculate rates using function
	successRate := calculatePercentage(successfulF, totalF)
	failureRate := calculatePercentage(failedF, totalF)

	// Print formatted output
	fmt.Println("\n" + header)
	fmt.Printf("Total Transactions: %d\n", total)
	fmt.Printf("Successful: %d\n", successful)
	fmt.Printf("Failed: %d\n", failed)
	fmt.Printf("Success Rate: %.0f%%\n", successRate)
	fmt.Printf("Failure Rate: %.0f%%\n", failureRate)
}
