package main

import "fmt"

// for -> only constructs in go for looping

func main(){
	// while loop

	// var i = 1 
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i +1 
	}

	// infinite loop

	// for {
	// 	println("1")
	// }

	// Classic for loop

	for i := 0; i<3; i++ {
		fmt.Println(i)

		// break  : stop the loop
		// continue : stop the current itteratiom
	}

	// Range

	for i := range 5 {
		fmt.Println(i)
	}
}