package main

import (
	"fmt"
	"time"
)

func main(){

	age := 5

	switch age {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default: // optional 
		fmt.Println("other")
	}

	// Multiple condition switch
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's workday")	
	}

	// Type switch
	whoAmI := func(i interface{}){
		switch t := i.(type){
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("Other", t)					
		}
	}

	whoAmI("golang")
}
