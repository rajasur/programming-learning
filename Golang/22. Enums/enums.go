package main

import "fmt"

// enumerated types

type orderStatus int

// type OrderStatus string

// const (
// 	Received  OrderStatus = "received"
// 	Confirmed             = "confirmed"
// 	Prepared              = "prepared"
// 	Delivered             = "delivered"
// )

const (
	Received orderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status orderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {
	changeOrderStatus(Prepared)
}
