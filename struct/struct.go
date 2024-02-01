package main

import (
	"fmt"
)

// Creating a struct
type student struct {
	star []int
}

// Push a new value into the stack
func (set *student) Push(num int) {
	set.star = append(set.star, num)
	fmt.Println("The value added :", num)
}

// Pop() function
func (set *student) Pop() int {
	if len(set.star) == 0 {
		fmt.Println("Stack is empty")
	} else {
		num := set.star[len(set.star)-1]
		set.star = set.star[:len(set.star)-1]
		fmt.Println("Pop value :", num)
	}
	return 0
}

// Top() function
func (set *student) Top() int {
	if len(set.star) == 0 {
		fmt.Println("Stack is empty")
	} else {
		value := set.star[len(set.star)-1]
		fmt.Println("top functionalities :", value)
	}
	return 0
}

// Main function
func main() {
	slice := make([]int, 0)
	stack := &student {
		star: slice,
	}
	var choice, num int
Loop:
	for {
		fmt.Println("press 1 - For performing the push")
		fmt.Println("press 2 - For performing the pop")
		fmt.Println("press 3 - For top functionalities")
		fmt.Println("press 4 - Display the values")
		fmt.Println("press 5 - To exit")
		fmt.Println("Enter your choice :")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("Enter the values to push")
			fmt.Scanln(&num)
			stack.Push(num)
		case 2:
			stack.Pop()
		case 3:
			stack.Top()
		case 4:
			if len(stack.star) == 0 {
				fmt.Println("Stack is empty", stack.star)
			} else {
				fmt.Println(stack.star)
			}
		case 5:
			fmt.Println("You are exited")
			break Loop
		default:
			fmt.Println("Invalid input")
		}
	}
}
