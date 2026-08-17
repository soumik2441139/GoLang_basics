package main

import "fmt"

//enumerated types

type OrderStatus int

const (
	Received OrderStatus = iota
	confirmed
	shipped
	delivered
)

func changeorderstatus(status OrderStatus) {
	fmt.Println("Order status changed to", status)
}

func main() {

	changeorderstatus(confirmed)

}
