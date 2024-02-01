package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name        string
	Age         int
	Active      bool
	LastLoginAt string
}

type Person struct {
	Name string
	Age  int
}

func main() {
	user1 := User{Name: "Nirmal", Age: 21, Active: true, LastLoginAt: "Today"}
	j, err := json.Marshal(user1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
	var a string = `{"Name":"Nirmal","Age":21}`
	var per = &Person{}
	err := json.Unmarshal([]byte(a), per)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshal:\nName: %s\nAge: %d", per.Name, per.Age)
}
