package main

import (
	"fmt"
)

func main() {
	// Arrays
	var marks [3]int

	marks[0] = 10
	marks[1] = 20
	marks[2] = 30

	fmt.Println(marks)

	num := [5]int{1, 2, 3, 4, 5}

	fmt.Println(len(num))

	// slices (most common collection type)
	res := []string{"Go", "java"}
	fmt.Println(len(res))

	var num1 []int // empty slice then append
	num1 = append(num1, 10)

	// length and capacity of slice
	// if we exceeding the capacity go grows the backing array (usually doubles first times)

	scores := make([]int, 0, 5)
	fmt.Println(scores, len(scores), cap(scores))

	scores = append(scores, 100, 200,300,400,500, 600,3,4,5,6,1,30,34)
	fmt.Println(scores, len(scores), cap(scores))

	// spread in slice

	todos := []int{1,2,3}
	more := []int{4,5}

	todos = append(todos, more...)

	// for range over slice
	views := []int{10,20,30,40,50}

	total := 0

	for i,v := range views {
		fmt.Println("index",i, "views",v)
		total = total + v
	} 
	fmt.Println("total",total)
}
