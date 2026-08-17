package main

import "fmt"

type payment struct {
	gateway paymenter
}
type paymenter interface {
	pay(amount float32)
}

// open close principle
func (p payment) makePayment(amount float32) {
	//razorpayPaymentGw := razorpay{}
	//stripePaymentGw := stripe{}
	//razorpayPaymentGw.pay(amount)
	//stripePaymentGw.pay(amount)

	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay:", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe:", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal:", amount)
}

func main() {
	/*stripePaymentGw := stripe{}
	newpayment := payment{
		gateway: stripePaymentGw,
	}*/
	paypalGw := paypal{}
	newpayment := payment{
		gateway: paypalGw,
	}
	newpayment.makePayment(100)
}
