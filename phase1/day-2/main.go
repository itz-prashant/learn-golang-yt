package main

import (
	"fmt"
	"time"
)

func main() {
	// if else
	age := 19

	if age >= 18 {
		fmt.Println("Eligible for vote")
	} else {
		fmt.Println("Not eligible for vote")
	}

	if age >= 18 {
		fmt.Println("Adult")
	} else if age >= 13 && age < 18 {
		fmt.Println("Teen")
	} else {
		fmt.Println("Kids")
	}

	// Switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	case time.Friday:
		fmt.Println("Halfday")
	default:
		fmt.Println("workinng day")
	}

	// Loops

	for i := 1; i <= 5; i++ {
		switch i {
		case 2:
			fmt.Println("Skip the value of i", i)
			continue
		default:
			fmt.Println("value of i is", i)
		}
	}

	// Loops on range

	nums := []int{1, 2, 3, 4, 5}

	for _, i := range nums {
		switch i {
		case 2:
			fmt.Println("Skip the value of i", i)
			continue
		default:
			fmt.Println("value of i is", i)
		}
	}
}
