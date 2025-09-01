package main

import (
	"fmt"
)

// // Sending function
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		println("Processing Number: ", num)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// // receiving function
// func additionNumbers(result chan int, num1 int, num2 int) {
// 	addition := num1 + num2
// 	result <- addition
// }

// // Goroutine Synchronizer
// func task(done chan bool) {
// 	defer func() {
// 		done <- true
// 	}()
// 	println("Processing...")
// }

// // Email sender messaging queue
// func emailSender(emailChan chan string, done chan bool) {

// 	defer func() { done <- true }()

// 	for email := range emailChan {
// 		fmt.Println("sending email to", email)
// 		time.Sleep(time.Second)
// 	}

// }

// // Email sender messaging queue type safety
// func emailSender(emailChan <-chan string, done chan<- bool) {

// 	defer func() { done <- true }()

// 	// emailChan <- "something@gmail.com" // data can't send
// 	//<-done // data can't receive

// 	for email := range emailChan {
// 		fmt.Println("sending email to", email)
// 		time.Sleep(time.Second)
// 	}

// }

func main() {
	// numChan := make(chan int)

	// // sending data through channel
	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// // receiving data

	// result := make(chan int)
	// go additionNumbers(result, 5, 6)

	// res := <-result
	// println("Addition of two numbers is:", res)

	// done := make(chan bool)
	// go task(done)
	// <-done // blocked until task is done

	// emailChan := make(chan string, 100) // define buffered channel which is never blocked

	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 1; i <= 5; i++ {

	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("Sending done...")

	// close(emailChan) // It is important to close the channel otherwise it will lead to a deadlock
	// <-done

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)

		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}

}
