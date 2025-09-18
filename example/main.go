package main

import (
	"fmt"
	"log"

	"github.com/jukbot/bathtext-go"
)

func main() {
	// Example usage of bathtext library
	amounts := []float64{0, 1.25, 100.50, 1000}

	fmt.Println("bathtext-go Example Usage")
	fmt.Println("=========================")
	fmt.Printf("Library Version: %s\n\n", bathtext.Version)

	for _, amount := range amounts {
		result, err := bathtext.Convert(amount)
		if err != nil {
			log.Printf("Error converting %.2f: %v", amount, err)
			continue
		}
		fmt.Printf("%.2f -> %s\n", amount, result)
	}

	// Example with error handling
	fmt.Println("\nError handling example:")
	_, err := bathtext.Convert(-10)
	if err != nil {
		fmt.Printf("Expected error for negative amount: %v\n", err)
	}
}