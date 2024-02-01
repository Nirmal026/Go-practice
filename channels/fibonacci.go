package main

import (
	"fmt"
)

// main function
func main() {
	// creating a channels
	number := make(chan int)
	end := make(chan bool)
	var value int
	// printing the input
	fmt.Println("Input :")
	fmt.Scanln(&value)
	// creating anonymous goroutine
	go func(value int) {
		fmt.Println("Output :")
		for i := 0; i < value; i++ {
			// read from channel
			fmt.Println(<-number)
		}
		end <- false
	}(value)
	fibonacci(number, end)
}

// function fibonacci
func fibonacci(number chan int, end chan bool) {
	x, y := 0, 1
	for {
		select {
		// write to channel
		case number <- x:
			x, y = y, x+y
		case <-end:
			return
		}
	}
}
