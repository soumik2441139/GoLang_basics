package main

import "fmt"

// by value
/*func changenum(num int) {
	num = 5
	fmt.Println("In changeNum", num)
}*/

// by reference
func changenum(num *int) {
	*num = 5
	fmt.Println("In changeNum", *num)
}

func main() {

	num := 1

	changenum(&num)

	// changenum(num)
	//fmt.Println("Memory address", &num)
	fmt.Println("After changeNum in main", num)
}
