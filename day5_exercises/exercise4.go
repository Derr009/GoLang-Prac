package main

import (
	"encoding/csv"
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

func parseEmployeesFromCSV(filePath string) ([]Employee, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var employees []Employee
	for i, record := range records {
		if i == 0 {
			continue // Skip header
		}

		salary, _ := strconv.ParseFloat(record[3], 64)

		emp := Employee{
			ID:         record[0],
			Name:       record[1],
			Department: record[2],
			Salary:     salary,
		}
		employees = append(employees, emp)
	}

	return employees, nil
}

func main() {
	employees, err := parseEmployeesFromCSV("Employees.csv")
	if err != nil {
		log.Fatalf("Error parsing CSV: %v", err)
	}

	fmt.Printf("Parsed %d employees into struct slice:\n", len(employees))
	for _, emp := range employees {
		fmt.Printf("%+v\n", emp)
	}
}
