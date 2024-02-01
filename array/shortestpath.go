package main

import (
	"fmt"
)

func main() {
	var size, sum int
	fmt.Println("Enter the length of the array :")
	_, err := fmt.Scanln(&size)
	if err != nil {
		panic(err)
	}
	var arr = make([]int, size)
	fmt.Println("Provide the array elements:")
	for i:=0;i<len(arr);i++ {
		_, err := fmt.Scanln(&arr[i])
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("The Given array is :",arr)
	fmt.Println("Provide the Target :")
	_, err = fmt.Scanln(&sum)
		if err != nil {
			panic(err)
		}
		var check int
	for i:=0; i<=int(size-1); i++ {
		for j:=0; j < i; j++ {
			if arr[i]+arr[j] == sum && check == 0 {
				fmt.Println("Shortest path : \n", j, ",", i, "-", arr[j], ",", arr[i])
				check = 1
			} else if arr[i]+arr[j] == sum {
				if check == 1 {
					fmt.Println("Other possibilities :")
				}
				fmt.Println(j, ",", i, "-", arr[j], ",", arr[i])
				check = 2
			}
		}
	}
}
