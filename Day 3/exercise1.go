package main

import "fmt"

func main() {
	numbers := []int{100, 250, 400, 125, 800}

	total := 0
	maxVal := numbers[0]
	var greaterThan300 []int

	for _, val := range numbers {
		total += val

		if val > maxVal {
			maxVal = val
		}

		if val > 300 {
			greaterThan300 = append(greaterThan300, val)
		}
	}

	avg := float64(total) / float64(len(numbers))

	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Average: %.2f\n", avg)
	fmt.Printf("Maximum: %d\n", maxVal)
	fmt.Printf("Values > 300: %v\n", greaterThan300)
}
