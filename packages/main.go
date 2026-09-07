package main

import (
	"awesomeProject/packages/auth"
	"awesomeProject/packages/user"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("admin", "password")

	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Email:    "admin@example.com",
		Username: "admin",
	}
	fmt.Println(user)

	color.Red(user.Email)

}
