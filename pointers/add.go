package main

import (
	"fmt"
)
// main function
func main() {
	var firstNumber, secondNumber, thirdNumber int
	fmt.Println("Input the first number :")
	fmt.Scanln(&firstNumber)
	fmt.Println("Input the second number :")
	fmt.Scanln(&secondNumber)
	thirdNumber = add(&firstNumber, &secondNumber)
	// Printing the output
	fmt.Printf("The sum of %d and %d is %d", firstNumber, secondNumber, thirdNumber)
}
func add(first *int, second *int) (third int) {
	third = *first + *second
	return third
}
