package main

import "fmt"

func sum (nums ...int)int{
	total := 0

	for _,num := range nums {
		total = total + num
	}

	return total
}

func main(){
	fmt.Println(2,3,4,5,"hello") // n number of elemnet pass in Println function

	result := sum(3,4,5,6,7,8,9)
	fmt.Println(result)

	// pass number using slice slice
	nums :=[]int{3,4,5,6,7}
	result2 := sum(nums...)
	fmt.Println(result2)
}