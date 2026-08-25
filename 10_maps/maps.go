package main

import (
	"fmt"
	"maps"
)

func main(){
	// creating map 
	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// get an element
	fmt.Println(m["name"], m["area"])
	fmt.Println(m["phone"]) // If key does not exist in the map then it returns zerp value


	a := make(map[string]int)

	a["age"] = 12
	a["price"] = 12

	fmt.Println(a["age"])
	fmt.Println(len(a))

	// delete elemnet
	delete(a, "price")

	// delete all element fo map
	clear(a)

	// another way 
	k := map[string]int{"price": 40, "phones": 3}
	println(k)

	// check element from map
	_ , ok := k["price"]

	if ok {
		fmt.Printf("all ok")
	}else{
		fmt.Printf("Not ok")
	}

	// check equal in map
	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 3}

	fmt.Println(maps.Equal(m1,m2))
}