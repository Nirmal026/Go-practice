package main

import (
	"fmt"
)

func main() {

	// Initializing the string
	var person = map[string]int{"nirmal": 1, "age": 21, "id": 53}
	fmt.Println(person)
	
	// Empty map declaration
	var student = map[string]int{}
	fmt.Println(student)

	// Map declaration using Make function
	employee := make(map[string]int)
	employee["nirmal"] = 10
	employee["prasad"] = 20
	fmt.Println(employee)

	// Accessing items in map
	personDetails := person["nirmal"]
	fmt.Println("i am", personDetails)

	// Adding the value in map
	sample := map[string]int{"name": 1, "age": 2, "id": 3}
	fmt.Println(sample)

	// Updating the value in map
	pet := map[string]string{"home": "dog", "house": "cat"}
	fmt.Println("before update :", pet)
	pet["home"] = "tiger"
	fmt.Println("after update :", pet)

	// Deleting the value from map
	animal := map[string]string{"forest": "lion", "town": "frog"}
	fmt.Println("before delete :", animal)
	delete(animal, "town")
	fmt.Println("after delete :", animal)

	// Iterate over a  map
	city := make(map[int]string)
	city[0] = "kovilpatti"
	city[1] = "satur"
	city[2] = "selam"
	for key, value := range city {
		fmt.Println("position :", key, "\nvalue :", value)
	}

	// Merge maps
	rollno := map[int]string{1: "suresh", 2: "kumar", 3: "surya"}
	rollnum := map[int]string{4: "gopal", 5: "vijay", 6: "vikram"}
	for stu, emp := range rollnum {
		rollno[stu] = emp
	}
	fmt.Println("Merged output :", rollno)

	var district = map[string]int{"orange": 4, "apple": 2, "cocanut": 1, "banana": 3}
	country := make([]string, 0, len(district))
	for k := range district {
		country = append(country, k)
		for x := range country {
			y := x + 1
			for y = range country {
				if country[x] < country[y] {
					country[x], country[y] = country[y], country[x]
				}
			}
		}
	}
	for _, k := range country {
		fmt.Println("sorted by keys :",k,district[k])
	}
}
