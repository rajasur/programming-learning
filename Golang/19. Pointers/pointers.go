package main

import "fmt"

func afterChange(number *int) {
	*number = 5
	fmt.Println("Value of number in afterChange function is:", *number)
}

func main() {
	number := 1
	afterChange(&number)                                        // passing the address of number
	fmt.Println("Value of number in main function is:", number) // value of number is changed to 5
}
