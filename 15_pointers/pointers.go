package main

import "fmt"

func changeNum(num *int) {
	*num = 5
	fmt.Println("In change num", *num)
}

func main() {
	var num = 1
	// fmt.Println("Memory address", &num)
	changeNum(&num)

	fmt.Println("After chnagenum in main", num)
}
