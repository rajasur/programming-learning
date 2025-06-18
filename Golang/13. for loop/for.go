package main

import "fmt"

func main() {
	// while loop using for
	i := 1
	for i <= 3 {
		fmt.Println("Value of i is:", i)
		i++
	}
	// infinite loop
	/*for {
		fmt.Println("This is an infinite loop")
	}*/
	// classic for loop
	for i := 0; i < 3; i++ {
		fmt.Println("Value of i in classic for loop is:", i)
	}
	for j := 0; j < 3; j++ {
		if j == 2 {
			continue
		}
		fmt.Println("Value of i in classic for loop is:", i)
	}

	// 1.22 range loop
	for i := range 3 {
		fmt.Println("Value of i in range loop is:", i)
	}
}
