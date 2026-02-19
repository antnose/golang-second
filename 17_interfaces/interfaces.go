package main

import "fmt"

type payment struct {
}

func (p payment) makePayment(amount float64) {
	razorPayPaymentGw := razorPay{}
	razorPayPaymentGw.pay(amount)
}

type razorPay struct {
}

func (r razorPay) pay(amount float64) {
	fmt.Println("Making payment using razorpay", amount)
}

func main() {
	newPayment := payment{}
	newPayment.makePayment(230)
}
