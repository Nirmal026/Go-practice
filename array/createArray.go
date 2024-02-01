package main

import (
	"fmt"
)

func main() {
	arr := [10]string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
	fmt.Println("Array of strings :")
	for i := 0; i < 10; i++ {
		fmt.Println(arr[i])
	}
}

