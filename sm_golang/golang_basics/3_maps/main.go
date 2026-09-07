package main

import "fmt"

func main() {
	// map[keyType]valueType
	marks := map[string]int{
		"math" : 40,
		"eng" : 30,
	}

	fmt.Println(marks["math"], len(marks))

	// create a empty map
	// make([key]valye)

	var scores map[string]int // nil map

	scores = make(map[string]int)

	scores["hindi"] = 40

	fmt.Println(scores)

	// delete key from map
	users := map[string]string{
		"u1" : "John",
		"u2" : "doe",
		"u3" : "demo", 
	}
	fmt.Println(users)
	delete(users, "u2")
	fmt.Println(users)

	// read value ok
	points := map[string]int{
		"a" :10,
		"b" : 0, // valid value
	}

	fmt.Println("a", points["a"])
	fmt.Println("b", points["b"])
	fmt.Println("c", points["c"])

	valB, okB := points["b"]
	valc, okc := points["c"]
	fmt.Println(valB, okB)
	fmt.Println(valc, okc)

	if val, ok := points["c"]; ok {
		fmt.Println(val,"is present")
	}

	prices := map[string]int{
		"a" : 100,
		"b" : 200,
	}

	for item, price := range prices{
		fmt.Println("item", item , "of", "price", price)
	}
}