package main

import "fmt"

func main() {
	fmt.Println("--- All Transactions ---")
	for i := 1; i <= 20; i++ {
		fmt.Printf("Transaction %d\n", i)
	}

	fmt.Println("\n--- Even Transactions Only ---")
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Printf("Transaction %d\n", i)
		}
	}
}
