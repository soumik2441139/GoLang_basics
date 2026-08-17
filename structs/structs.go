package main

import (
	"fmt"
	"time"
)

// embedding
type customer struct {
	name  string
	phone string
}

// composition
type order struct {
	// Define fields for the order struct
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
	customer
}

func newOrder(id string, amount float32, status string) *order {
	// initial setup
	myorder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myorder

}

func (o *order) changestatus(status string) {
	o.status = status
}

func (o order) getamount() float32 {
	return o.amount
}

func main() {

	//myorder := newOrder("12345", 100.50, "pending")
	//fmt.Println("myorder struct", myorder)

	neworder := order{
		id:     "1",
		amount: 30,
		status: "received",
	}
	newcustomer := customer{
		name:  "Soumik",
		phone: "1234567890",
	}
	neworder.customer = newcustomer
	fmt.Println(neworder.customer)

	//inline struct
	/*language := struct {
		name   string
		isgood bool
	}{"golang", true}
	fmt.Println(language)*/

	// if you don't set any field , default value will be set for that field
	/*order1 := order{
		id:        "12345",
		amount:    100.50,
		status:    "pending",
		createdAt: time.Now(),
	}
	order1.changestatus("completed")
	fmt.Println("order struct", order1)
	fmt.Println("order amount", order1.getamount())*/
	/*order2 := order{
		id:        "12346",
		amount:    200.75,
		status:    "completed",
		createdAt: time.Now(),
	}
	fmt.Println("order2 struct", order2)

	order1.createdAt = time.Now()
	//fmt.Println(order1.status)
	fmt.Println("order struct", order1)*/
}
