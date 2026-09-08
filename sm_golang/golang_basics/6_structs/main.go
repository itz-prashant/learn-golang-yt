package main

import "fmt"

// struct groups related fields into one type
type user struct {
	id int
	name string
	email string
	age int
}

// method pointer reciever
func (u *user) changeName(name string) {
	u.name = name
}

func (u user) getEmail() string{
	return  u.email
}

func main() {
	u1 := user{
		id: 1,
		name: "Prashant",
		email: "p@gmail.com",
		age: 27,
	}
	fmt.Println(u1)
	u1.changeName("Update")
	fmt.Println(u1)

	fmt.Println(u1.getEmail())
}