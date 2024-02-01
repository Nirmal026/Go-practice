package main

import "fmt"

// Function twice for twicing the values
func twice(sum chan int) {
	for i := 1; i <= 5; i++ {
		sum <- i * 2
	}
	close(sum)
}

// Main functiom
func main() {
	// Creating channels
	numbers := make(chan int)
	// calling goroutine
	go twice(numbers)
	// Using Foreach loop
	for res := range numbers {
		fmt.Println(res)
	}
}
