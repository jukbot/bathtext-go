// Package bathtext provides functionality to convert numbers to Thai currency full text.
package bathtext

import (
	"errors"
	"fmt"
)

// Version of the bathtext library
const Version = "1.0.0"

// ErrInvalidInput is returned when the input is invalid
var ErrInvalidInput = errors.New("invalid input")

// Convert converts a number to Thai currency full text.
// This is a placeholder implementation that will be replaced with actual logic.
func Convert(amount float64) (string, error) {
	if amount < 0 {
		return "", fmt.Errorf("%w: negative amounts not supported", ErrInvalidInput)
	}

	// Placeholder implementation
	// TODO: Implement actual Thai currency text conversion
	return fmt.Sprintf("%.2f บาท", amount), nil
}

// ConvertBaht converts a number to Thai currency full text in baht.
// This is a convenience function that calls Convert.
func ConvertBaht(amount float64) (string, error) {
	return Convert(amount)
}