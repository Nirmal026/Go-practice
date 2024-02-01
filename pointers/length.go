package main

import (
	"bufio"
	"fmt"
	"os"
)
// main function
func main() {
	var letter int
	fmt.Println("Input a string :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	alpabets := scanner.Text()
	alpabet := []rune(alpabets)
	letter = length(&alpabet)
	fmt.Printf("The length of the given string %s is : %d", alpabets, letter)
}

func length(word *[]rune) (letter int) {
	letter = len(*word)
	return letter
}
