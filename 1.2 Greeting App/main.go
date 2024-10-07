package main

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now()
	ft := dt.Format("2006-01-02 15:04:05")
	fmt.Println("Current date and time is: ", ft)

	hr := dt.Hour()

	if hr >= 6 && hr < 11 {
		fmt.Println("Good Morning!")
	} else if hr >= 11 && hr < 16 {
		fmt.Println("Good Afternoon!")
	} else if hr >= 16 && hr < 21 {
		fmt.Println("Good evening!")
	} else {
		fmt.Println("Good night!")
	}

}
