package main

import (
	"fmt"
)
// Employee struct
type Employee struct {
	name        string
	dateofbirth int
	ID          int
}
// main function
func main() {
	emp := Employee{"nirmal", 2000, 53}
	point := &emp
	fmt.Println((*point).name)
	fmt.Println((*point).dateofbirth)
	fmt.Println((*point).ID)
}
