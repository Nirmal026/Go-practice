package main

import (
	"fmt"
)

type office struct{
	empname string
	empid int
	address string
}

func main(){
	      var name string
		  fmt.Print("Enter your name :")
		  fmt.Scanln(&name)
	      var age int
		  fmt.Print("Enter your age :")
		  fmt.Scanf("%d",&age)
	      var phone int64= 123456789
		  var percentage float32= 96.5
		  var boolvalue bool=true
		  var arr = [5]string{"Html","Css","Js","Golang","MySql"}
		  var slice = arr[1:2]
		  fmt.Println("My Slice :",slice)
		  fmt.Printf("I'm %s\nMy Age is %d\nPhone Number is %d\nMy percentage is %g\nIt is %t\nMy fav courses %q\n",name,age,phone,percentage,boolvalue,arr)
		  fmt.Println(office{"Nirmal Prasad",53,"Thoothukudi"})
		  var city string="kovilpatti"
		  fmt.Println(city)
		  var employee byte ='#'
		  fmt.Printf("%d\n",employee)
		  var student rune ='汉'
		  fmt.Printf("%d",student)
}