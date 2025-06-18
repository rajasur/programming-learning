package main

func add(x int, y int) int {
	return x + y
}
func subtract(x int, y int) int {
	return x - y
}
func multiply(x int, y int) int {
	return x * y
}
func divide(x int, y int) int {
	if y == 0 {
		return 0 // Handle division by zero
	}
	return x / y
}

func getLanguages() (string, string, string, string, string) {
	// This function returns a list of programming languages
	// as a tuple. The languages are hardcoded for demonstration purposes.
	return "golang", "json", "python", "java", "c++"
}

func processIt(fn func(a int) int) {
	fn(1)
}

func processIt2() func(a int) int {
	// This function returns another function that takes an integer and returns an integer.
	return func(a int) int {
		return a * 2
	}
}

// processIt2 demonstrates how to return a function from another function.
func main() {
	// Example usage of the functions
	a := 10
	b := 5

	sum := add(a, b)
	diff := subtract(a, b)
	prod := multiply(a, b)
	quot := divide(a, b)
	lang1, lang2, lang3, lang4, lang5 := getLanguages()
	println("Languages:", lang1, lang2, lang3, lang4, lang5)

	fn := func(a int) int {
		return 2
	}

	fn2 := processIt2()
	println(fn2) // Using the returned function to process a value
	processIt(fn)
	println("Sum:", sum)
	println("Difference:", diff)
	println("Product:", prod)
	println("Quotient:", quot)
}
