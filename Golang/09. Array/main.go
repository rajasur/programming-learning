package main

import "fmt"

func main() {

	var nums [4]int
	fmt.Println("Length of nums array is:", len(nums)) // Length of array is 4
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	nums[3] = 40
	fmt.Println("Elements of nums array are:", nums) // Elements of nums array are: [10 20 30 40]

	arr := []int{1, 2, 3, 4, 5}
	result := sum(arr)
	fmt.Println("Sum of array elements is:", result)
}
