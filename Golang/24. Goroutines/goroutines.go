package main

import (
	"fmt"
	"time"
)

// func task(id int) {
// 	fmt.Println("doing task", id)
// }

func main() {
	for i := 0; i <= 10; i++ {
		// go task(i) // Start a goroutine for each task

		// closure function with goroutine
		go func(i int) {
			fmt.Println("doing task", i)
		}(i) // Use an anonymous function to capture the loop variable
	}

	time.Sleep(time.Second * 2) // Wait for goroutines to finish
}
