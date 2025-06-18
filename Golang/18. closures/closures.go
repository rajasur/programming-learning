package main

import "fmt"

func counter() func() int {
	var count int
	return func() int {
		count += 1
		return count
	}
}

func main() {
	res := counter() // create a new counter
	fmt.Println(res())
	fmt.Println(res())
}
