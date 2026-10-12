package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("Employees.csv")
	if err != nil {
		log.Fatalf("Failed to open input file: %v", err)
	}
	defer inputFile.Close()

	reader := csv.NewReader(inputFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV: %v", err)
	}

	outputFile, err := os.Create("High_salary.csv")
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Write header
	writer.Write(records[0])

	count := 0
	for i, record := range records {
		if i == 0 {
			continue
		}

		salary, err := strconv.ParseFloat(record[3], 64)
		if err == nil && salary > 100000 {
			writer.Write(record)
			count++
		}
	}

	fmt.Printf("Successfully written %d high-salary records to High_salary.csv\n", count)
}
