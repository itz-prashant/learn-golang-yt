package main

import (
	"fmt"
	"github.com/itz-prashant/podcast/auth"
	"github.com/itz-prashant/podcast/user"
)

func main(){
	auth.LoginWithCredentials("a@gmail.com", "secret")
	session := auth.GetSession()
	fmt.Println("session",session)

	user := user.User{
		Email: "user@email.com",
		Name: "John Doe",
	}

	fmt.Println("User", user.Email, user.Name)
}