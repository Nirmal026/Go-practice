package main

import "fmt"

// Main function
func main() {
	var arr = make([]int, 7)
	var num int
	fmt.Printf("Input the first element of the array : \n")
	fmt.Scanln(&num)
	arr[0] = num
	fmt.Println("Array elements :")
	fmt.Printf("array_nums[0] = %d\n",num)
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i-1] * 2
		fmt.Printf("array_nums[%d] = %d\n",i,arr[i])
	}
}
