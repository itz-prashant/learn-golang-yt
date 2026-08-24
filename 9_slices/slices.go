package main

import (
	"fmt"
	"slices"
)

// slices -> dynamic
// most used construct in go
// + useful methods

func main(){

	// uninitialized slice is nil
	var nums []int

	fmt.Println(nums) // [] -> this is nil

	var age = make([]int, 0,5)
	fmt.Println(age) // [0 0] -> initial size

	// fmt.Println(cap(age)) // capacity -> maximum numbers of element can fit

	age = append(age, 1,2,3,4)
	fmt.Println(age)
	fmt.Println(cap(age))

	// direct declare slice
	score := []int{2}
	score[0] = 1
	score = append(score, 2,3)
	fmt.Println(score)
	fmt.Println(cap(score))

	// Copy fucntion

	var age2 = make([]int, len(age))
	copy(age2, age)

	fmt.Println(age2)

	// Slice operator
	var student = []int{1,2,3}
	fmt.Println(student[0:2]) // [1,2]
	fmt.Println(student[:2]) // [1,2]
	fmt.Println(student[1:]) // [2,3]

	// slice package
	var nums1 = []int{1,2}
	var nums2 = []int{1,2}

	fmt.Println(slices.Equal(nums1, nums2)) // return bool

	// 2d slices

	num3 := [][]int{{1,2,3},{4,5,6}}
		fmt.Println(num3)
}