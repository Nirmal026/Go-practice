package main

import (
	"fmt"
)

// Main function
func main() {
	var sum int
	// Getting the input from the user
	fmt.Println("Input any number to find sum of digits :")
	fmt.Scanln(&sum)
	// Calling the function
	var number = recursion(sum)
	// Printing the output
	fmt.Printf("The Sum of digits of %d = %d", sum, number)
}
func recursion(num int) int {
	if num == 0 {
		return 0
	} else {
		return ((num % 10) + recursion(num/10))
	}
}
