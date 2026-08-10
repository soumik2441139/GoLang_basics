package main

import "fmt"

func main() {

	/*simple switch

	i := 5

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}*/

	/*switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}*/

	whoami := func(i interface{}) {
		switch t := i.(type) {
		case int:
			println("I am an int")
		case string:
			println("I am a string")
		case bool:
			fmt.Println("Its a boolean")
		default:
			println("Its an unknown type", t)
		}
	}
	whoami(42)
	whoami("hello")
	whoami(3.14)
}
