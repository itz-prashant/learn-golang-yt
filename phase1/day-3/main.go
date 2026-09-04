package main

import "fmt"

type person struct {
	name string
	age  int
}

type rect struct {
	width, height int
}

func changeNum(num *int) {
	*num = 5
}

// value reciever
func (r rect) area () int {
	return r.width * r.height
}

func (r *rect) scale(factor int) {
	r.width *= factor
	r.height *= factor
}

func main() {
	// Pointers
	x := 10

	p := &x

	fmt.Println(p)
	fmt.Println(*p)

	*p = 20
	x = 30

	fmt.Println("new value of x", x)
	fmt.Println("new value of p", *p)

	num := 1
	changeNum(&num)
	fmt.Println(num)

	// Structs
	p1 := person{
		name: "Prashant",
		age: 28,
	}

	p2 := &p1

	fmt.Println("Before change",p1)
	p2.name = "Kumar"
	fmt.Println("After change",p1)

	r := rect{10,15}

	fmt.Println("Area", r.area())
	r.scale(2)
	fmt.Println("scaled area", r.area())
	fmt.Println("r", r)
}
