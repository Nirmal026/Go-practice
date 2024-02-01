package main

import (
	"bufio"
	"fmt"
	"os"
)

// Main function
func main() {
	var words string
	fmt.Println("Input any string :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	alpabets := scanner.Text()
	alpabet := []rune(alpabets)
	// Calling the function
	reverse(alpabet, &words)
	// Printing the output
	fmt.Println("The reversed string is :", words)
}
func reverse(alpabet []rune, words *string) {
	*words += string(alpabet[len(alpabet)-1])
	// fmt.Println(*words)
	alpabet = alpabet[:len(alpabet)-1]
	// fmt.Println(alpabet)
	if len(alpabet) > 0 {
		reverse(alpabet, words)
	}
}
