package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time
}

func newOrder(id string, amount float64, status string) *Order {
	myOrder := Order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

func (o *Order) changeStatus(status string) {
	o.status = status
}

func main() {
	order := Order{
		id:     "1",
		amount: 40.500,
		status: "received",
	}

	order.createdAt = time.Now()

	fmt.Println(order)

	order.changeStatus("confirmed")
	fmt.Println("My order", order)
}
