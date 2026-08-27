package main

import "fmt"

// enumeratde types
type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status OrderStatus){
	fmt.Println(status)
}

func main() {
	changeOrderStatus(Confirmed)
}