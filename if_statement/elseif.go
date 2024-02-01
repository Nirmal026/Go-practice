package main

import (
	"fmt"
)

var amount, hundred, fifty, twenty, ten, five, two, one int
// Main function
func main() {
	// Input amount from user
	fmt.Print("In the amount :")
	_, err := fmt.Scanln(&amount)
	if err != nil {
		fmt.Println("Invalid Input")
		return
	}else if amount <= 0 {
		fmt.Println("Invalid Input")
		return
	}
	for {
		if amount >= 100 {
			hundred = amount / 100
			amount = amount - (hundred * 100)
		} else if amount >= 50 {
			fifty = amount / 50
			amount = amount - (fifty * 50)
		} else if amount >= 20 {
			twenty = amount / 20
			amount = amount - (twenty * 20)
		} else if amount >= 10 {
			ten = amount / 10
			amount = amount - (ten * 10)
		} else if amount >= 5 {
			five = amount / 5
			amount = amount - (five * 5)
		} else if amount >= 2 {
			two = amount / 2
			amount = amount - (two * 2)
		} else if amount >= 1 {
			one = amount / 1
			amount = amount - (one * 1)
		} else{
			break
		}
	}
	// print required notes
	fmt.Println(hundred, "Note(s) of 100.00")
	fmt.Println(fifty, "Note(s) of 50.00")
	fmt.Println(twenty, "Note(s) of 20.00")
	fmt.Println(ten, "Note(s) of 10.00")
	fmt.Println(five, "Note(s) of 5.00")
	fmt.Println(two, "Note(s) of 2.00")
	fmt.Println(one, "Note(s) of 1.00")
}
