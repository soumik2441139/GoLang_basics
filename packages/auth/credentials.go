package auth

import "fmt"

func LoginWithCredentials(username, password string) bool {
	fmt.Println("Logging in with username:", username, "and password:", password)
	if username == "admin" && password == "password" {
		return true
	}
	return false
}
