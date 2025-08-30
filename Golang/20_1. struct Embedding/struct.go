package main

import (
	"fmt"
	"time"
)

type customer struct {
	name        string
	phoneNumber int
}
type order struct {
	id        int
	status    string
	country   string
	recivedAt time.Time
	customer
}

func main() {
	myOrder := order{
		id:      1,
		status:  "received",
		country: "India",
	}

	fmt.Println("My order is: ", myOrder)

}
