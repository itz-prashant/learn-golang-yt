package main

import "fmt"

const age = 30

func main() {
	// Can't reasign
	const name string = "golang"

	fmt.Printf(name)

	fmt.Println(age)

	// Multiple constant

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
