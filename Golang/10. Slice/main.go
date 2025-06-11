package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Println("Slice:", slice)
	slice = append(slice, 4, 5)
	fmt.Println("After appending:", slice)

	// Declaration of slice with make
	slice2 := make([]int, 3, 5)
	fmt.Println("Length of Slice2 is:", len(slice2))
	fmt.Println("Capacity of Slice2 is:", cap(slice2))

}
