package main

import "fmt"

func main(){
	age := 18

	if age >= 18 {
		fmt.Println("Person is an adult")
	}else {
		fmt.Println("Is not an adult")
	}

	// Else if

	if age >=18 {
		fmt.Println("Person is an adult")
	}else if age >= 12 {
		fmt.Println("Person is teenager")
	}else {
		fmt.Println("Person is kid")
	}

	var role string = "admin"

	var hasPermission bool = true

	if role == "admin" && hasPermission {
		fmt.Println("Yes")
	}

	// Varibale declare directly inside if constructure

	if age := 15; age >= 18 {
		fmt.Println("Person is an adult", age)
	}else if age >= 12 {
		fmt.Println("Person is teenager")
	}else {
		fmt.Println("Person is kid")
	}

}
