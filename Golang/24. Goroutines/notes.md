# Goroutines in Go

Goroutines are lightweight threads managed by the Go runtime. They enable concurrent programming, allowing multiple functions to run independently.

## Key Concepts

### 1. What is a Goroutine?
- A function executing independently, concurrently with other goroutines.
- Created using the `go` keyword.

### 2. Creating a Goroutine

```go
go functionName(args)
```

**Example:**
```go
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from goroutine!")
}

func main() {
    go sayHello()
    time.Sleep(1 * time.Second) // Wait for goroutine to finish
    fmt.Println("Main function ends")
}
```

### 3. Main Goroutine
- The `main` function runs in its own goroutine.
- The program exits when the main goroutine finishes, even if other goroutines are running.

### 4. Anonymous Goroutines

```go
go func() {
    fmt.Println("Anonymous goroutine")
}()
```

### 5. Passing Arguments

```go
go printMessage("Hello", 3)

func printMessage(msg string, count int) {
    for i := 0; i < count; i++ {
        fmt.Println(msg)
    }
}
```

### 6. Synchronization

#### WaitGroup

```go
import "sync"

var wg sync.WaitGroup

wg.Add(1)
go func() {
    defer wg.Done()
    fmt.Println("Goroutine with WaitGroup")
}()
wg.Wait()
```

#### Channels

```go
ch := make(chan string)
go func() {
    ch <- "Message from goroutine"
}()
msg := <-ch
fmt.Println(msg)
```

### 7. Communication Between Goroutines

- Use **channels** to send and receive data safely.

### 8. Buffered Channels

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
fmt.Println(<-ch)
fmt.Println(<-ch)
```

### 9. Closing Channels

```go
close(ch)
```

### 10. Select Statement

- Waits on multiple channel operations.

```go
select {
case msg1 := <-ch1:
    fmt.Println(msg1)
case msg2 := <-ch2:
    fmt.Println(msg2)
default:
    fmt.Println("No communication")
}
```

### 11. Goroutine Leaks

- Occur when goroutines are blocked forever (e.g., waiting on a channel that never receives data).

### 12. Best Practices

- Always synchronize goroutines (WaitGroup, channels).
- Avoid sharing data without synchronization (use mutexes if needed).
- Limit the number of goroutines to avoid resource exhaustion.

## Summary

- Goroutines are fundamental for concurrency in Go.
- Use channels and synchronization primitives to coordinate goroutines.
- Proper management is crucial to avoid leaks and race conditions.
