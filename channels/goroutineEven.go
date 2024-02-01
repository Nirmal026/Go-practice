package main

import (
	"fmt"
)

// main function
func main() {
	var size int
	fmt.Println("Enter the size :")
	fmt.Scanln(&size)
	// creating a slice
	digits := make([]int, size)
	// creating a channels
	oddNum := make(chan int, size)
	evenNum := make(chan int, size)
	fmt.Println("Input:")
	for i := 0; i < len(digits); i++ {
		fmt.Scanln(&digits[i])
	}
	fmt.Println("Output:")
	// calling goroutine
	go operation(digits, oddNum, evenNum)
	// printing the ODD numbers
	for i := range oddNum {
		fmt.Println("ODD : ", i)
	}
	// printing the EVEN numbers
	for j := range evenNum {
		fmt.Println("EVEN: ", j)
	}
}

// function operation for finding ODD and EVEN numbers
func operation(digits []int, oddNum chan int, evenNum chan int) {
	for _, value := range digits {
		if value%2 != 0 {
			oddNum <- value
		} else {
			evenNum <- value
		}
	}
	// closing the channels
	close(oddNum)
	close(evenNum)
}
