package main

import "fmt"

// main function
func main() {
	var lines, numbers int
	len := 1
	fmt.Print("Input number of lines: ")
	fmt.Scan(&lines)
	fmt.Print("Number of characters in a line: ")
	fmt.Scan(&numbers)
	fmt.Println("Output :")
	for i := 1; i <= lines; i++ {
		for j := 1; j <= numbers; j++ {
			fmt.Print(len, " ")
			len++
		}
		fmt.Print("\n")
	}
}
