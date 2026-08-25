package main

import "fmt"

// iterating over data structure
func main(){
	nums := []int{4,5,6}

	for i :=0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// sum using range
	sum := 0
	for i, num := range nums {
		fmt.Println(i) // print index
		sum  = sum + num
		fmt.Println(num)  // 4 5 6
	}
	fmt.Println(sum)

	m := map[string]string{"firstName":"John", "lastName":"doe"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// use ramge in string
	// i -> starting byte of rune

	for i, c := range "golang"{
		fmt.Println(i, c) // c = unicode (code point rule)
		fmt.Println(i, string(c)) // print character

	}
}