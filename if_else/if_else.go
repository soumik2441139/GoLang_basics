package main

import "fmt"

func main() {
	age := 16
	if age >= 18 {
		fmt.Println("You are eligible to vote.")
	} else {
		println("You are not eligible to vote.")
	}

	var role = "admin"
	var hasPermission = true
	if role == "admin" && hasPermission {
		fmt.Println("You have admin access.")
	} else {
		fmt.Println("You do not have admin access.")
	}

	if agee := 15; agee >= 18 {
		fmt.Println("person is an adult", agee)
	} else if agee >= 12 {
		println("person is a teenager", agee)
	}

	//go does not have ternary operator, you will have to use normal if else

}
