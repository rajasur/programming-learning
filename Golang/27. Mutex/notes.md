# Golang Mutex: Detailed Notes

## What is a Mutex?

A **Mutex** (short for "mutual exclusion") is a synchronization primitive used to prevent multiple goroutines from accessing shared resources (like variables, data structures, or files) at the same time. It ensures that only one goroutine can access the critical section of code at a time, preventing race conditions.

## Why Use Mutex?

- To protect shared data from concurrent access.
- To avoid race conditions and ensure data consistency.
- To coordinate goroutines in concurrent programming.

## Mutex in Go

Go provides the `sync.Mutex` type in the `sync` package.

### Basic Usage

```go
import (
    "fmt"
    "sync"
)

var (
    counter int
    mu      sync.Mutex
)

func increment() {
    mu.Lock()         // Acquire the lock
    counter++
    mu.Unlock()       // Release the lock
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            increment()
        }()
    }
    wg.Wait()
    fmt.Println("Counter:", counter)
}
```

**Explanation:**
- `mu.Lock()` blocks if another goroutine holds the lock.
- `mu.Unlock()` releases the lock.
- Always unlock after locking, ideally using `defer`.

## Important Points

- **Lock/Unlock:** Always pair `Lock()` with `Unlock()`. Use `defer` to ensure unlock even if a panic occurs.
- **Deadlocks:** If a goroutine tries to lock a mutex it already holds, it will deadlock.
- **Zero Value:** The zero value of `sync.Mutex` is usable (no need to initialize).
- **Not Copyable:** Mutexes must not be copied after first use.

## Example: Protecting a Map

```go
type SafeMap struct {
    mu sync.Mutex
    m  map[string]int
}

func (s *SafeMap) Write(key string, value int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.m[key] = value
}

func (s *SafeMap) Read(key string) int {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.m[key]
}
```

## sync.RWMutex

- `sync.RWMutex` allows multiple readers or one writer.
- Use `RLock()`/`RUnlock()` for reading, `Lock()`/`Unlock()` for writing.

```go
var rw sync.RWMutex

rw.RLock()   // Multiple readers allowed
rw.RUnlock()

rw.Lock()    // Only one writer allowed
rw.Unlock()
```

## Common Pitfalls

- Forgetting to unlock: leads to deadlocks.
- Locking too much: reduces concurrency.
- Copying mutexes: leads to undefined behavior.

## When Not to Use Mutex

- For communication between goroutines, prefer channels.
- Use mutexes for protecting shared state, not for signaling.

## Summary

- Mutexes are essential for safe concurrent access to shared resources.
- Always lock before accessing shared data and unlock after.
- Use `sync.Mutex` for exclusive access, `sync.RWMutex` for multiple readers.
- Avoid common mistakes like forgetting to unlock or copying mutexes.
