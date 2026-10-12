package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Employee struct {
	ID         string  `json:"employee_id"`
	Name       string  `json:"name"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

func main() {
	file, err := os.Open("Employees.csv")
	if err != nil {
		log.Fatalf("Failed to open CSV: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV: %v", err)
	}

	var employees []Employee
	for i, record := range records {
		if i == 0 {
			continue
		}

		salary, _ := strconv.ParseFloat(record[3], 64)
		employees = append(employees, Employee{
			ID:         record[0],
			Name:       record[1],
			Department: record[2],
			Salary:     salary,
		})
	}

	// Convert slice of structs to formatted JSON
	jsonData, err := json.MarshalIndent(employees, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	// Write JSON to file
	err = os.WriteFile("Employees.json", jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing JSON file: %v", err)
	}

	fmt.Println("Successfully converted CSV records to Employees.json")
}
