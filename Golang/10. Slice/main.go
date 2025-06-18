package main

import "fmt"

func main() {
	var nums []int           // Declare an empty slice of integers
	fmt.Println(nums == nil) // Check if the slice is nil

	slice := []int{1, 2, 3}
	fmt.Println("Slice:", slice)
	slice = append(slice, 4, 5)
	fmt.Println("After appending:", slice)

	// Declaration of slice with make
	slice2 := make([]int, 3, 5)
	fmt.Println("Length of Slice2 is:", len(slice2))
	fmt.Println("Capacity of Slice2 is:", cap(slice2))

	var slice3 = make([]int, 0, 5)
	slice3 = append(slice3, 1, 2, 3, 4, 5)
	fmt.Println("Slice3:", slice3)
	slice3 = append(slice3, 6, 7, 8) // Appending more elements
	fmt.Println("After appending more elements:", slice3)
	fmt.Println("Length of Slice3 is:", len(slice3))
	fmt.Println("Capacity of Slice3 is:", cap(slice3))

	num := []int{}
	fmt.Println("Length of nums slice is:", len(num))   // Length of slice is 0
	fmt.Println("Capacity of nums slice is:", cap(num)) // Capacity of slice is 0
	nums = append(nums, 1, 2, 3, 4, 5)
	fmt.Println("After appending elements:", num)                       // After appending elements: [1 2 3 4 5]
	fmt.Println("Length of nums slice after appending is:", len(num))   // Length of slice after appending is 5
	fmt.Println("Capacity of nums slice after appending is:", cap(num)) // Capacity of slice after appending is 5
	fmt.Println("Is nums slice nil?", num == nil)                       // Is nums slice nil? false
	fmt.Println("Is nums slice empty?", len(num) == 0)                  // Is nums slice empty? false
	fmt.Println("Is nums slice non-empty?", len(num) > 0)               // Is nums slice non-empty? true
	fmt.Println("Is nums slice non-empty?", len(num) > 0)               // Is nums slice non-empty? true
	fmt.Println("Is nums slice empty?", len(num) == 0)                  // Is nums slice empty? false
	fmt.Println("Is nums slice nil?", num == nil)                       // Is nums slice nil? false
	fmt.Println("Is nums slice non-empty?", len(num) > 0)               // Is nums slice non-empty? true
	fmt.Println("Is nums slice empty?", len(num) == 0)                  // Is nums slice empty? false
	fmt.Println("Is nums slice nil?", num == nil)                       // Is nums slice nil? false
	fmt.Println("Is nums slice non-empty?", len(num) > 0)               // Is nums slice non-empty? true
	fmt.Println("Is nums slice empty?", len(num) == 0)                  // Is nums slice empty? false
	fmt.Println("Is nums slice nil?", num == nil)

	var slice4 = make([]int, 2, 5)
	slice4[0] = 10
	slice4[1] = 20
	fmt.Println("Slice4:", slice4) // Slice4: [10 20]

}
