package main

import "fmt"

type stack[T int | string] struct {
	elements []T
}

func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	nums := []int{1, 2, 3}

	names := []string{"golang", "java", "c++"}
	printSlice(nums)
	printSlice(names)

	myStack := stack[int]{
		elements: []int{1, 2, 3},
	}

	fmt.Println(myStack)
}
