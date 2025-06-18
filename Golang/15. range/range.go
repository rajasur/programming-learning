package main

import (
	"fmt"
)

func main() {
	// iterating over data structures using range
	nums := []int{1, 2, 3, 4, 5}

	for i, v := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	sum := 0
	for _, v := range nums {
		sum = sum + v
	}
	fmt.Println("Sum of all elements:", sum)

	// map iteration
	m := map[string]string{
		"fname": "John",
		"lanme": "Doe",
	}

	for k, v := range m {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}

	for k := range m {
		fmt.Printf("Key: %s\n", k)
	}

	// string iteration
	for i, c := range "golang" {
		fmt.Println(i, string(c))                      // i is the index and c is the rune (unicode point)
		fmt.Printf("Index: %d, Character: %c\n", i, c) // c is unicode point and i is starting byte of rune
	}
}
