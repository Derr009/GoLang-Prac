package validator

import (
	"errors"
	"fmt"
	"strings"
)

// ValidateID checks if the transaction/entity ID is valid
func ValidateID(id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("ID cannot be empty")
	}
	if len(id) < 3 {
		return fmt.Errorf("ID '%s' is too short (minimum 3 characters required)", id)
	}
	return nil
}

// ValidateAmount ensures the transaction amount is positive
func ValidateAmount(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be greater than zero, got %.2f", amount)
	}
	return nil
}

// ValidateStatus checks against allowed lifecycle status strings
func ValidateStatus(status string) error {
	s := strings.ToUpper(strings.TrimSpace(status))
	switch s {
	case "SUCCESS", "FAILED", "CANCELLED", "PENDING":
		return nil
	default:
		return fmt.Errorf("invalid status '%s' (must be SUCCESS, FAILED, CANCELLED, or PENDING)", status)
	}
}

// ValidateCustomer checks if the customer identifier is valid
func ValidateCustomer(customer string) error {
	if strings.TrimSpace(customer) == "" {
		return errors.New("customer name/ID cannot be empty")
	}
	return nil
}

// ValidateRecord runs all validators on a set of transaction fields
func ValidateRecord(id string, amount float64, status string, customer string) []error {
	var errs []error

	if err := ValidateID(id); err != nil {
		errs = append(errs, err)
	}
	if err := ValidateAmount(amount); err != nil {
		errs = append(errs, err)
	}
	if err := ValidateStatus(status); err != nil {
		errs = append(errs, err)
	}
	if err := ValidateCustomer(customer); err != nil {
		errs = append(errs, err)
	}

	return errs
}
