package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Println("Enter the size :")
	fmt.Scanln(&size)
	digits := make([]int, size)
	primeNum := make(chan int, size)
	fmt.Println("Input:")
	for i := 0; i < len(digits); i++ {
		fmt.Scanln(&digits[i])
	}
	fmt.Println("Output:")
	go primeFilter(digits, primeNum)
	for i := range primeNum {
		fmt.Println("Prime : ", i)
	}
}

func primeFilter(digits []int, primeNum chan int) {
	for i, value := range digits {
		if i%value != 0 {
			primeNum <- value
		}
	}
	fmt.Println(<-primeNum)
	close(primeNum)
}
