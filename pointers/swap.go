package main

import (
	"fmt"
)
// main function
func main() {
	var firstNumber, secondNumber, thirdNumber int
	fmt.Println("Input the value of 1st element :")
	fmt.Scanln(&firstNumber)
	fmt.Println("Input the value of 2st element :")
	fmt.Scanln(&secondNumber)
	fmt.Println("Input the value of 3st element :")
	fmt.Scanln(&thirdNumber)
	// Print statement before swapping
	fmt.Printf("The value before swapping are :\n element 1 = %d\n element 2 = %d\n element 3 = %d\n", firstNumber, secondNumber, thirdNumber)
	swap(&firstNumber, &secondNumber, &thirdNumber)
	// Print statement after swapping
	fmt.Printf("The value after swapping are :\n element 1 = %d\n element 2 = %d\n element 3 = %d\n", firstNumber, secondNumber, thirdNumber)
}
func swap(first *int, second *int, third *int) {
	var temp int
	temp = *second
	*second = *first
	*first = *third
	*third = temp
}
