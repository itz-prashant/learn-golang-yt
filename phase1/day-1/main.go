// variables
// Three method to declare variable

package main

import "fmt"

const salary = 22

func main() {
	// 1st method
	var name string
	name = "golang"

	var age float32
	age = 28.3
	fmt.Println(name, age)

	// 2nd method
	var num = 12
	fmt.Println(num)

	// 3rd method 
	student := "John doe"
	fmt.Println(student)

	// get data type
	lang := "golang"
	fmt.Printf("value %v and data_type is %T \n", lang, lang)

	fmt.Println(salary)

}
