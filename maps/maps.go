package main

import (
	"fmt"
	"maps"
)

func main() {

	//m := make(map[string]string)

	//m["name"] = "Soumik"
	//m["lastname"] = "Manna"
	//fmt.Println(m["name"], m["lastname"])
	//IMP: if key does not exist, it will return zero value of that type

	/*m := make(map[string]int)
	m["age"] = 30
	m["price"] = 50
	fmt.Println(len(m))

	delete(m, "price")
	//clear(m)
	fmt.Println(m)*/
	//fmt.Println(m)
	/*m := map[string]int{"age": 30, "price": 50}

	k, ok := m["price"]
	fmt.Println(k)
	if !ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}*/

	m1 := map[string]int{"age": 30, "price": 50}
	m2 := map[string]int{"age": 30, "price": 50}

	fmt.Println(maps.Equal(m1, m2))

}
