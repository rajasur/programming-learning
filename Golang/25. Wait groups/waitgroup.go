package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done() // it executes at the end
	fmt.Println("doing task", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go task(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
