package main

import "fmt"

func add(a , b int) int {
	return  a + b
}

func sumAndProduct(a, b int) (int, int) {
	sum := a + b
	product := a * b

	return  sum , product
}

// named return value
func divide (a int, b int) (john int) {
	john = a/b
	return 
}

// variadic function
func sumAll(nums ...int) int {
	fmt.Println("nums",nums)
	total := 0

	for _ ,currentvalue := range nums {
		total = total + currentvalue
	}

	return total
}

func main() {
	res := add(10, 20)
	fmt.Println(res)

	sum, product := sumAndProduct(20,30)

	fmt.Println(sum, product)

	fmt.Println(divide(20, 2))

	total := sumAll(1,2,3,4,5)
	fmt.Println(total)

	values := []int{10,23}
	fmt.Println(sumAll(values...))

}