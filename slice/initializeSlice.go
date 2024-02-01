package main

import (
	"fmt"
	// "sort"
)

// Main function
func main() {
	// Initialize Slice with values
	var numbers = []int{1, 2, 3, 4}
	var letters = []string{"one", "two", "three", "four"}
	fmt.Printf("%v\n", numbers)
	fmt.Printf("%v\n", letters)

	// Declare Slice using new Keyword
	var num = new([50]int)[0:10]
	fmt.Println(num)

	// Accessing the Slice
	// You can access the Slice by using Index value
	fmt.Println("Accessing the Slice :", numbers[0])
	fmt.Println("Accessing the Slice :", letters[2])

	// Updating the value in Slice
	letters[3] = "five"
	fmt.Println(letters)

	// Remove item from slice
	letters = append(letters[0:2])
	fmt.Println("nirmal",letters)
	letters = append(letters[:4])
	fmt.Println("niru",letters)

	// Copy a Slice
	var animal = []string{"tiger", "lion", "wolf"}
	pet := make([]string,len(animal))
	c := copy(pet, animal)
	fmt.Println(c, "Elements copied")
	fmt.Println("Animals :", animal)
	fmt.Println("Pets :", pet)

	// Loop Through a Slice
	loop := []string{"car", "bus", "bike"}
	for Index, element := range loop {
		fmt.Println(Index, element)
	}

	// Reverse the slice
	places := []string{"Chennai", "kovilpatti","salem", "madurai", "thoothukudi"}
	// fmt.Println("Original Format :", places)
	// sort.Sort(sort.Reverse(sort.StringSlice(places)))
	// fmt.Println("Reversed Format :", places)

     //using for loop
	 fmt.Println("The original string is :",places)
	 for i, j := 0, len(places)-1; i < j; i, j = i+1, j-1 {
		places[i], places[j] = places[j], places[i]
	}
	fmt.Println("The reversed string is :",places)
	 


}