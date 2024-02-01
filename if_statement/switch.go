package main

import (
	"fmt"
)

// Main function
func main() {
	var amount int
	// Input amount from user
	fmt.Print("In the amount :")
	fmt.Scanln(&amount)
	switch {
	case amount >= 100:
		var notes int = amount / 100
		// print required note
		fmt.Printf("%d Note(s) of 100.00\n", notes)
		amount = amount - (notes * 100)
		fallthrough
	case amount >= 50:
		var notes int = amount / 50
		// print required note
		fmt.Printf("%d Note(s) of 50.00\n", notes)
		amount = amount - (notes * 50)
		fallthrough
	case amount >= 20:
		var notes int = amount / 20
		// print required note
		fmt.Printf("%d Note(s) of 20.00\n", notes)
		amount = amount - (notes * 20)
		fallthrough
	case amount >= 10:
		var notes int = amount / 10
		// print required note
		fmt.Printf("%d Note(s) of 10.00\n", notes)
		amount = amount - (notes * 10)
		fallthrough
	case amount >= 5:
		var notes int = amount / 5
		// print required note
		fmt.Printf("%d Note(s) of 5.00\n", notes)
		amount = amount - (notes * 5)
		fallthrough
	case amount >= 2:
		var notes int = amount / 2
		// print required note
		fmt.Printf("%d Note(s) of 2.00\n", notes)
		amount = amount - (notes * 2)
		fallthrough
	case amount >= 1:
		var notes int = amount / 1
		// print required note
		fmt.Printf("%d Note(s) of 1.00\n", notes)
		amount = amount - (notes * 1)
	default:
		fmt.Println("Invalid Input")
	}
}
