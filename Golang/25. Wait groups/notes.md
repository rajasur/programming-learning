# Wait Groups in Go

## What are Wait Groups?

Wait Groups are a synchronization primitive in Go that allow you to wait for a collection of goroutines to finish executing. They are part of the `sync` package and provide a way to coordinate multiple concurrent operations.

## When to Use Wait Groups

- When you need to wait for multiple goroutines to complete before proceeding
- When you have a known number of goroutines that need to finish
- When you want to ensure all concurrent tasks are done before program termination
- When coordinating parallel processing tasks

## Why Use Wait Groups?

- **Synchronization**: Ensures all goroutines complete before moving forward
- **Clean shutdown**: Prevents main function from exiting before goroutines finish
- **Resource management**: Helps avoid race conditions and incomplete operations
- **Deterministic behavior**: Makes concurrent programs more predictable

## How to Use Wait Groups

### Basic Structure

```go
package main

import (
        "fmt"
        "sync"
        "time"
)

func main() {
        var wg sync.WaitGroup
        
        // Add number of goroutines to wait for
        wg.Add(3)
        
        // Launch goroutines
        for i := 0; i < 3; i++ {
                go func(id int) {
                        defer wg.Done() // Signal completion
                        fmt.Printf("Goroutine %d working...\n", id)
                        time.Sleep(time.Second)
                        fmt.Printf("Goroutine %d done\n", id)
                }(i)
        }
        
        // Wait for all goroutines to complete
        wg.Wait()
        fmt.Println("All goroutines completed")
}
```

### Key Methods

- `Add(delta int)`: Adds delta to the WaitGroup counter
- `Done()`: Decrements the counter by one (equivalent to Add(-1))
- `Wait()`: Blocks until the counter becomes zero

### Best Practices

1. **Always call Done()**: Use `defer wg.Done()` at the beginning of goroutines
2. **Pass by pointer**: Pass WaitGroup as pointer to avoid copying
3. **Add before launching**: Call `Add()` before starting goroutines
4. **One Wait() call**: Typically call `Wait()` only once in main goroutine

### Common Patterns

```go
// Pattern 1: Fixed number of workers
func fixedWorkers() {
        var wg sync.WaitGroup
        numWorkers := 5
        
        wg.Add(numWorkers)
        for i := 0; i < numWorkers; i++ {
                go worker(i, &wg)
        }
        wg.Wait()
}

// Pattern 2: Dynamic number of tasks
func dynamicTasks(tasks []Task) {
        var wg sync.WaitGroup
        
        for _, task := range tasks {
                wg.Add(1)
                go func(t Task) {
                        defer wg.Done()
                        processTask(t)
                }(task)
        }
        wg.Wait()
}
```

### Common Pitfalls

- Forgetting to call `Done()`
- Calling `Add()` after `Wait()`
- Not using pointers when passing WaitGroups
- Race conditions between `Add()` and `Wait()`