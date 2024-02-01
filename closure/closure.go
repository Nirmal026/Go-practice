package main

import "fmt"


func testClosure() func() int  {
	c := 0
	return func() int {
		c += 1
		return c
	}
}

func main() {
	c := testClosure()
	fmt.Println("Output :", c())
}
