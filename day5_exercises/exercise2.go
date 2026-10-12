package main

import (
	"encoding/csv"
	"fmt"

	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("Employees.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV: %v", err)
	}

	fmt.Println("=== Employees with Salary > 100,000 ===")
	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}

		salary, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			continue
		}

		if salary > 100000 {
			fmt.Printf("ID: %s | Name: %s | Dept: %s | Salary: $%.2f\n", record[0], record[1], record[2], salary)
		}
	}
}
