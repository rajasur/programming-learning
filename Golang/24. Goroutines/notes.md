
# Goroutines in Go

## What are Goroutines?

Goroutines are lightweight threads managed by the Go runtime. They are functions or methods that run concurrently with other functions or methods. Goroutines are one of Go's key features for concurrent programming.

## Why Use Goroutines?

- **Lightweight**: Much lighter than OS threads (2KB initial stack vs 8MB for OS threads)
- **Efficient**: Go runtime multiplexes goroutines onto OS threads
- **Scalable**: Can create thousands of goroutines without significant overhead
- **Simple**: Easy syntax with the `go` keyword
- **Built-in concurrency**: Native support for concurrent programming

## When to Use Goroutines?

- I/O operations (file reading, network requests)
- Parallel processing of data
- Background tasks
- Web servers handling multiple requests
- Producer-consumer patterns
- Any task that can benefit from concurrent execution

## How to Create Goroutines

### Basic Syntax
```go
go functionName()
go func() {
    // anonymous function
}()
```

### Example 1: Simple Goroutine
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
    time.Sleep(1 * time.Second) // Wait for goroutine
    fmt.Println("Main function")
}
```

### Example 2: Multiple Goroutines
```go
func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    for i := 1; i <= 3; i++ {
        go worker(i)
    }
    time.Sleep(2 * time.Second)
}
```

## Goroutine Communication

### Using Channels
```go
func main() {
    ch := make(chan string)
    
    go func() {
        ch <- "Hello from goroutine"
    }()
    
    message := <-ch
    fmt.Println(message)
}
```

## Important Concepts

### Main Goroutine
- Program starts with the main goroutine
- When main goroutine exits, all other goroutines are terminated
- Use synchronization to wait for goroutines to complete

### Goroutine Scheduling
- Go runtime uses M:N scheduler
- M goroutines mapped to N OS threads
- Cooperative scheduling with preemption

## Best Practices

1. **Always synchronize**: Use channels, WaitGroups, or other sync primitives
2. **Avoid shared state**: Prefer message passing over shared memory
3. **Handle goroutine lifecycle**: Ensure proper cleanup
4. **Don't create too many**: Though lightweight, goroutines still consume resources

## Common Patterns

### Worker Pool
```go
func workerPool() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Start workers
    for w := 1; w <= 3; w++ {
        go worker(jobs, results)
    }
    
    // Send jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Collect results
    for a := 1; a <= 5; a++ {
        <-results
    }
}
```

### WaitGroup Pattern
```go
import "sync"

func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Goroutine %d\n", id)
        }(i)
    }
    
    wg.Wait()
}
```

## Key Takeaways

- Goroutines make concurrent programming simple and efficient
- Use `go` keyword to start a goroutine
- Always handle synchronization properly
- Goroutines are not threads but are multiplexed onto threads
- Communication through channels is preferred over shared variables
