package main

import (
	"fmt"
	"github.com/sj8687/my-go-app/auth"
	"github.com/sj8687/my-go-app/user"
	"github.com/fatih/color"

)

func main() {
  	auth.LoginWithCredentials("sj", "bsdk")
	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Email: "user@email.com",
		// Name:  "John Doe",
	}

	// fmt.Println(user.Email, user.Name)
	color.Green(user.Email)

}