package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("Employees.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read header row
	header, err := reader.Read()
	if err != nil {
		log.Fatalf("Failed to read header: %v", err)
	}
	fmt.Printf("Header: %v\n\n", header)

	// Read data rows line by line
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading CSV row: %v", err)
		}

		empID := record[0]
		name := record[1]
		dept := record[2]
		salary := record[3]

		fmt.Printf("ID: %s | Name: %s | Dept: %s | Salary: %s\n", empID, name, dept, salary)
	}
}
