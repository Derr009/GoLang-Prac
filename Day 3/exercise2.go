package main

import "fmt"

func main() {
	statusCount := map[string]int{
		"SUCCESS":   120,
		"FAILED":    30,
		"CANCELLED": 15,
	}

	fmt.Println("=== Status Counts ===")
	for status, count := range statusCount {
		fmt.Printf("%s: %d\n", status, count)
	}
}
