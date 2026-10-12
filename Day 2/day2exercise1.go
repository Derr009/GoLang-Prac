package main

import "fmt"

func calculateSuccessRate(successful, total float64) float64 {
	if total == 0 {
		return 0.0
	}
	return (successful / total) * 100
}

func calculateFailureRate(failed, total float64) float64 {
	if total == 0 {
		return 0.0
	}
	return (failed / total) * 100
}

func calculateAverage(totalAmount float64, count int) float64 {
	if count == 0 {
		return 0.0
	}
	return totalAmount / float64(count)
}

func main() {
	total := 50.0
	successful := 42.0
	failed := 8.0
	totalAmount := 10500.0

	fmt.Printf("Success Rate: %.2f%%\n", calculateSuccessRate(successful, total))
	fmt.Printf("Failure Rate: %.2f%%\n", calculateFailureRate(failed, total))
	fmt.Printf("Average Value: %.2f\n", calculateAverage(totalAmount, int(total)))
}
