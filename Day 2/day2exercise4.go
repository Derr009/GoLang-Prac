package main

import "fmt"

func classifyStatus(status string) string {
	switch status {
	case "SUCCESS":
		return "Transaction completed successfully."
	case "FAILED":
		return "Transaction failed due to processing error."
	case "CANCELLED":
		return "Transaction was cancelled by user."
	case "PENDING":
		return "Transaction is currently awaiting processing."
	default:
		return "Unknown transaction status."
	}
}

func main() {
	statuses := []string{"SUCCESS", "FAILED", "CANCELLED", "PENDING", "UNKNOWN"}

	for _, s := range statuses {
		fmt.Printf("[%s]: %s\n", s, classifyStatus(s))
	}
}
