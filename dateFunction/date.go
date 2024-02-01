package main

import (
	"fmt"
	"time"
)

func main() {
	// this function returns the present time
	currentTime := time.Now()
	currentTime.Date()
	var weekdays int
	fmt.Printf("Today's date : %d-%02d-%02d\n", currentTime.Year(), currentTime.Month(), currentTime.Day())
	// Printing the time
	fmt.Println("current Month in numeric :", currentTime.Format("01"))
	fmt.Println("current month in word :", currentTime.Format("January"))
	week := currentTime.Weekday()
	weeks := int(week)
	for i := 0; i < 7; i++ {
		if i+weeks == 3 {
			weekdays = i - 7
			day := currentTime.AddDate(0, 0, weekdays)
			fmt.Printf("last wednesday date : %d-%02d-%02d\n", day.Year(), day.Month(), day.Day())
		} else if weeks-i == 3 {
			weekdays = i - 3
			day := currentTime.AddDate(0, 0, weekdays)
			fmt.Printf("last wednesday date : %d-%02d-%02d\n", day.Year(), day.Month(), day.Day())
		}
	}
	lastDay := time.Now().AddDate(0, 1, -currentTime.Day())
	fmt.Printf("last day of month : %s\n", lastDay.Format("Monday"))
	for i := 1; i <= 5; i++ {
		firstDay := currentTime.AddDate(0, 0, -i)
		fmt.Printf("%02d-%s-%d - %s\n", firstDay.Day(), firstDay.Format("Jan"), firstDay.Year(), firstDay.Format("Monday"))
	}
}
