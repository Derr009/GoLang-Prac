package main

import "fmt"

func evaluateScore(score int) string {
	if score >= 90 {
		return "Excellent"
	} else if score >= 75 {
		return "Good"
	} else if score >= 60 {
		return "Average"
	} else {
		return "Needs Improvement"
	}
}

func main() {
	score := 82
	result := evaluateScore(score)
	fmt.Printf("Score: %d -> %s\n", score, result)
}
