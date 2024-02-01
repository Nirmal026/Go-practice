package main

import (
	"fmt"
)

func main() {
	var sec,hours,minutes,seconds int
	fmt.Println("Input seconds :")
	fmt.Scanln(&sec)
	hours = (sec/3600);
	minutes = (sec -(3600*hours))/60;
	seconds = (sec -(3600*hours)-(minutes*60));
	fmt.Printf("There are :\nH:M:S - %d:%d:%d\n",hours,minutes,seconds);
}