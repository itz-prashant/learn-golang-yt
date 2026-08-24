package main

import "fmt"

// numbered sequence of specific length
func main(){
	var nums [4]int // decleration

	fmt.Println((len(nums)))

	nums[0] = 1 // add new element on zero index

	fmt.Println(nums) // print all array [1,0,0,0]
	fmt.Println(nums[0]) // print single element of array

	var vals [4]bool
	fmt.Println(vals) // [all false]

	var name [5]string
	name[0] = "golang"
	fmt.Println(name) 

	// to declare it in single line
	age := [4]int{12,14,16,18}
	fmt.Println(age)

	// 2d array
	score := [2][2]int{{2,3},{5,6}}
	fmt.Println(score)

	// - fixed size , that is predictable
	// - memory optimazation
	// - content time access
}