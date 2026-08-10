package main

import "fmt"

// for-> only construct in go for looping
func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//while loop
	j := 1
	for j <= 3 {
		fmt.Println(j)
		j = j + 1
	}
	//infinite loop
	//for {
	//	println("1")
	//}

	//range loop
	for k := range 3 {
		fmt.Println(k)
	}
}
