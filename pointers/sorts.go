package main

import (
	"fmt"
)

// main function
func main() {
	var size int
	fmt.Print("Input the number of elements to store in the array :")
	_, err := fmt.Scanln(&size)
	if err != nil {
		panic(err)
	}
	var arr = make([]int, size)
	fmt.Printf("Input %d number of elements in the array :\n", size)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("element - %d :", i+1)
		_, err := fmt.Scanln(&arr[i])
		if err != nil {
			panic(err)
		}
	}
	sortarray(&arr)
	fmt.Printf("The elements in the array after sorting : \n")
	for i := 0; i < len(arr); i++ {
		// Printing the output
		fmt.Printf("Element - %d : %d \n", i+1, arr[i])
	}
}
func sortarray(a *[]int) {
	for i := 0; i < len(*a)-1; i++ {
		{
			if (*a)[i] > (*a)[i+1] {
				(*a)[i], (*a)[i+1] = (*a)[i+1], (*a)[i]
				i = -1
			}
		}
	}
}
