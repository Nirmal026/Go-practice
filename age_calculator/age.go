package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	var sonBirthDate, motherBirthDate string
	// Printing son's date of birth
	fmt.Println("your date of birth (DD-MM-YYYY) :")
	fmt.Scan(&sonBirthDate)
	sonBirthDay, err := time.Parse("2-1-2006", sonBirthDate)
	if err != nil {
		panic(err)
	}
	sonAge := int(currentTime.Sub(sonBirthDay).Hours() / 24 / 365)
	// Printing son's age
	fmt.Println("your current age is :", sonAge)
	// Printing mother's date of birth
	fmt.Println("your mother's date of birth (DD-MM-YYYY) :")
	fmt.Scan(&motherBirthDate)
	motherBirthDay, err := time.Parse("2-1-2006", motherBirthDate)
	if err != nil {
		panic(err)
	}
	motherAge := int(sonBirthDay.Sub(motherBirthDay).Hours() / 24 / 365)
	// Printing the age of the mother when you son born
	if motherAge < 18{
		fmt.Println("Your mother's date of birth is incorrect")
		return
	} else {
		fmt.Println("The age of your mother when you were born :", motherAge)
	}
	// The days left for your birthday
	birthDay := time.Date(currentTime.Year()+1, sonBirthDay.Month(), sonBirthDay.Day(), 0, 0, 0, 0, time.Local)
	birthDays := int(birthDay.Sub(currentTime).Hours()/24) + 1
	if birthDays < 365 {
		fmt.Println("Days left for your birthday :", birthDays)
	} else {
		birthDay := time.Date(currentTime.Year(), sonBirthDay.Month(), sonBirthDay.Day(), 0, 0, 0, 0, time.Local)
		birthDays := int(birthDay.Sub(currentTime).Hours()/24) + 1
		if birthDays == 0 {
			fmt.Println("Today is your birthday")
		} else {
			fmt.Println("Days left for your birthday :", birthDays)
		}
	}
}
