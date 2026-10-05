package main

import "fmt"

func main() {
	total := 5000.0
	successful := 4675.0
	failed := 325.0

	// Calculate percentages
	successPct := (successful / total) * 100
	failurePct := (failed / total) * 100

	fmt.Printf("Success %%: %.0f%%\n", successPct)
	fmt.Printf("Failure %%: %.0f%%\n", failurePct)
}
