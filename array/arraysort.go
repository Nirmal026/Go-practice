package main
 
import (
    "fmt"
    //"sort"
)

func main() {
    // arr := [...]int{8,3,2,1,3,2} //Using Sort function method
	// fmt.Println("Input :",arr) 
    // arrSort := arr[:]                        
    // sort.Ints(arrSort)
    // fmt.Println("Output :", arr)
    var size,temp int
    fmt.Println("Enter the size of the array :")
    _, err := fmt.Scanln(&size)
	if err != nil {
		panic(err)
	}
    var arr = make([]int, size) // using For loop method
    for i:=0;i<len(arr);i++{
		_, err := fmt.Scanln(&arr[i])
		if err != nil {
			panic(err)
		}
	}
    fmt.Println("Arrays :",arr)
    for i := 0; i < len(arr) - 1; i++{
        if arr[i] > arr[i+1] {
            temp = arr[i]
            arr[i] = arr[i+1]
            arr[i+1] = temp
            i = -1
        }
    }
    // Print the Output
    fmt.Println("Output :",arr)

}