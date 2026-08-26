package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func newOrder(id string, amount float32, status string) *order {
	// initial setup

	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

func main() {
	// if you don't set any field , default value is zero value
	// zero value -> int -> 0 , string -> "" , book -> false
	myOrder := order{
		id:     "1",
		amount: 347.7,
		status: "received",
	}

	myOrder.createdAt = time.Now()
	fmt.Println(myOrder)
	fmt.Println(myOrder.status)

	myOrder2 := order{
		id:        "2",
		amount:    450,
		status:    "received",
		createdAt: time.Now(),
	}

	myOrder2.changeStatus("confirmed")
	fmt.Println(myOrder2)
	fmt.Println(myOrder.getAmount())

	myOrder3 := newOrder("1", 230, "Paid")
	fmt.Println(myOrder3)

	// direct decalare

	lang := struct {
		name string
		isGood bool
	}{"goland", true} // same order

	fmt.Println(lang)
}
