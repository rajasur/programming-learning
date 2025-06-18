package main

func sum(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func sum2(nums ...interface{}) int {
	sum := 0
	for _, v := range nums {
		if num, ok := v.(int); ok {
			sum += num
		}
	}
	return sum
}
func main() {
	result := sum(1, 2, 3, 4, 5)
	println("Sum of numbers is:", result) // Sum of numbers is: 15
	result = sum(10, 20, 30)
	println("Sum of numbers is:", result) // Sum of numbers is: 60
	result = sum()                        // No arguments
	println("Sum of numbers is:", result) // Sum of numbers is: 0

	nums := []int{1, 2, 3, 4, 5}
	result = sum(nums...)                       // Using a slice with variadic function
	println("Sum of slice numbers is:", result) // Sum of slice numbers is: 15
	result = sum2(1, "two", 3.0, 4, 5)
	println("Sum of mixed numbers is:", result)                                   // Sum of mixed numbers is: 10
	println("Sum of mixed numbers with interface is:", sum2(1, "two", 3.0, 4, 5)) // Sum of mixed numbers with interface is: 10
}

// The sum function takes a variable number of integer arguments and returns their sum.
// The main function demonstrates the usage of the sum function with different sets of arguments.
// It prints the sum of the numbers passed to the sum function.
// The sum function uses the variadic parameter syntax (nums ...int) to accept a variable number of arguments.
// The for loop iterates over the nums slice, adding each value to the sum variable.
// The main function calls the sum function with different sets of integers and prints the results.
// The first call to sum passes five integers, the second call passes three integers, and the third call passes no arguments.
// The output will show the sum of the numbers for each call to the sum function.
// The sum function is a simple example of how to use variadic functions in Go.
// Variadic functions allow you to pass a variable number of arguments to a function.
// This is useful when you want to perform operations on a variable number of inputs without needing to define a fixed number of parameters.
// The main function serves as the entry point of the program, demonstrating the functionality of the sum function.
// The sum function is defined to take a variable number of integer arguments using the ellipsis (...) syntax.
// The function iterates over the provided integers, calculates their sum, and returns the result.
// The main function calls the sum function with different sets of integers and prints the results to the console.
// This code is a simple demonstration of how to use variadic functions in Go
