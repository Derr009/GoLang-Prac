package main

import "fmt"

func analyzeTransactions(amounts []float64) (float64, float64, float64, float64) {
	if len(amounts) == 0 {
		return 0, 0, 0, 0
	}

	total := 0.0
	highest := amounts[0]
	lowest := amounts[0]

	for _, amount := range amounts {
		total += amount
		if amount > highest {
			highest = amount
		}
		if amount < lowest {
			lowest = amount
		}
	}

	avg := total / float64(len(amounts))
	return total, avg, highest, lowest
}

func main() {
	amounts := []float64{100, 250, 500, 1200, 50}

	total, avg, highest, lowest := analyzeTransactions(amounts)

	fmt.Printf("Total: %.2f\n", total)
	fmt.Printf("Average: %.2f\n", avg)
	fmt.Printf("Highest: %.2f\n", highest)
	fmt.Printf("Lowest: %.2f\n", lowest)
}
