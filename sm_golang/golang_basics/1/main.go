package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// var declaration
	var name string
	name = "go"
	fmt.Println(name)
	fmt.Println(strings.ToUpper(name))

	var num int = 64
	fmt.Println(num)

	fmt.Println(math.Sqrt(25)) // 5

	// Short declaration
	likes1, comment := 3000, "hello"
	fmt.Println(likes1, comment)

	views1 := 1000
	views2 := 2000
	totalViews := views1 + views2

	likes := 20
	likes++
	likes++

	avgViews := totalViews / 2

	fmt.Println(totalViews, likes, avgViews)
}
