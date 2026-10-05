package main

import "fmt"

func main() {
	// Input variables
	employeeID := "E101"
	employeeName := "Rahul"
	projects := 4
	hoursWorked := 168

	// Calculate average hours per project
	avgHours := hoursWorked / projects

	// Output
	fmt.Printf("Employee: %s\n", employeeID)
	fmt.Printf("Name: %s\n", employeeName)
	fmt.Printf("Projects: %d\n", projects)
	fmt.Printf("Hours: %d\n", hoursWorked)
	fmt.Printf("Average hours/project: %d\n", avgHours)
}
