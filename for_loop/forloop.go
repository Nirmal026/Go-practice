package main

import (
	"fmt"
)

// Main function
func main() {
	var number, prime int
	var count bool
	// Input from user
	fmt.Println("Input a number: ")
	_, err := fmt.Scanf("%d", &number)
	for i := 2; i <= number; i++ {
		count = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				count = false
				break
			}
		}
		if count {
			prime++
		}
	}
	switch {
	// Printing error messages
	case err != nil:
		fmt.Println("Invalid Input")
	case number <= 0:
		fmt.Println("Input error")
	default:
		// Printing the Output
		fmt.Printf("Number of prime numbers which are less than or equal to %d \n%d", number, prime)
	}
}
