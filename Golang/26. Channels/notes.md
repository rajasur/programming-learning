
# Go Channels - Complete Guide

## What are Channels?

Channels are a fundamental feature in Go that enable goroutines to communicate with each other and synchronize their execution. They are typed conduits through which you can send and receive values between goroutines.

## Why Use Channels?

- **Goroutine Communication**: Enable safe data exchange between concurrent goroutines
- **Synchronization**: Coordinate execution and prevent race conditions
- **CSP Model**: Implement Communicating Sequential Processes paradigm
- **Memory Safety**: Avoid shared memory issues through message passing
- **Deadlock Prevention**: Built-in mechanisms to detect and prevent deadlocks

## When to Use Channels?

- Passing data between goroutines
- Coordinating concurrent operations
- Implementing producer-consumer patterns
- Building pipeline architectures
- Synchronizing goroutine execution
- Broadcasting signals across multiple goroutines

## Channel Types

### Unbuffered Channels
```go
ch := make(chan int)        // Synchronous channel
ch <- 42                    // Send (blocks until received)
value := <-ch              // Receive (blocks until sent)
```

### Buffered Channels
```go
ch := make(chan int, 5)    // Buffer size 5
ch <- 1                    // Non-blocking until buffer full
ch <- 2
value := <-ch             // Non-blocking until buffer empty
```

### Directional Channels
```go
// Send-only channel
func sender(ch chan<- int) {
    ch <- 42
}

// Receive-only channel
func receiver(ch <-chan int) {
    value := <-ch
}
```

## Channel Operations

### Creating Channels
```go
ch1 := make(chan int)           // Unbuffered
ch2 := make(chan string, 10)    // Buffered
ch3 := make(chan bool)          // Boolean channel
```

### Sending and Receiving
```go
ch <- value        // Send
value := <-ch      // Receive
value, ok := <-ch  // Receive with status check
```

### Closing Channels
```go
close(ch)

// Check if channel is closed
value, ok := <-ch
if !ok {
    // Channel is closed
}
```

### Range Over Channels
```go
for value := range ch {
    fmt.Println(value)
}
// Loop exits when channel is closed
```

## Select Statement

Handle multiple channel operations:

```go
select {
case msg1 := <-ch1:
    fmt.Println("Received from ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received from ch2:", msg2)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout")
default:
    fmt.Println("No channels ready")
}
```

## Common Patterns

### Worker Pool
```go
func workerPool(jobs <-chan int, results chan<- int) {
    for job := range jobs {
        results <- job * 2
    }
}
```

### Fan-out/Fan-in
```go
// Fan-out: distribute work
func fanOut(in <-chan int, out1, out2 chan<- int) {
    for val := range in {
        select {
        case out1 <- val:
        case out2 <- val:
        }
    }
}
```

### Pipeline
```go
func stage1(in <-chan int, out chan<- int) {
    for val := range in {
        out <- val * 2
    }
    close(out)
}
```

## Best Practices

1. **Channel Ownership**: The goroutine that creates a channel should close it
2. **Avoid Goroutine Leaks**: Always ensure goroutines can exit
3. **Use Buffered Channels Carefully**: Can hide synchronization issues
4. **Don't Close Receive-only Channels**: Compile-time error
5. **Check Channel Status**: Use the two-value receive when needed

## Common Pitfalls

- **Deadlocks**: Sending/receiving on blocked channels
- **Goroutine Leaks**: Goroutines waiting forever on channels
- **Race Conditions**: Multiple goroutines accessing shared state
- **Closing Closed Channels**: Causes panic
- **Sending on Closed Channels**: Causes panic

## Example: Complete Channel Usage

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Create channels
    jobs := make(chan int, 5)
    results := make(chan int, 5)
    done := make(chan bool)

    // Worker goroutine
    go func() {
        for job := range jobs {
            results <- job * job
        }
        close(results)
        done <- true
    }()

    // Send jobs
    for i := 1; i <= 5; i++ {
        jobs <- i
    }
    close(jobs)

    // Collect results
    go func() {
        for result := range results {
            fmt.Println("Result:", result)
        }
    }()

    // Wait for completion
    <-done
    time.Sleep(100 * time.Millisecond) // Allow results to print
}
```
