package main

import "fmt"

type shape interface {
	area() float64
	parameter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) parameter() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) parameter() float64 {
	return 2 * 3.14 * c.radius
}

func main() {
	var s shape

	s = rect{10,5}
	fmt.Println("Area of rect", s.area())
	fmt.Println("parameter of rect", s.parameter())

	s = circle{10}
	fmt.Println("Area of circle", s.area())
	fmt.Println("parameter of circle", s.parameter())
}
