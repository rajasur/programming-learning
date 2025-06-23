# Go Wait Groups

A **WaitGroup** in Go is used to wait for a collection of goroutines to finish executing. It is part of the `sync` package.

## Key Concepts

- **Purpose:** Synchronize concurrent goroutines.
- **Type:** `sync.WaitGroup`
- **Main Methods:**
    - `Add(int)`: Sets the number of goroutines to wait for.
    - `Done()`: Decrements the counter by one (called by each goroutine when finished).
    - `Wait()`: Blocks until the counter is zero.

## Basic Usage

```go
package main

import (
        "fmt"
        "sync"
)

func worker(id int, wg *sync.WaitGroup) {
        defer wg.Done() // Signal completion
        fmt.Printf("Worker %d starting\n", id)
        // Simulate work
        fmt.Printf("Worker %d done\n", id)
}

func main() {
        var wg sync.WaitGroup

        for i := 1; i <= 3; i++ {
                wg.Add(1) // Increment counter
                go worker(i, &wg)
        }

        wg.Wait() // Wait for all workers to finish
        fmt.Println("All workers done")
}
```

## Important Points

- **Call `Add` before starting goroutines.**
- **Call `Done` exactly once per goroutine.**
- **`Wait` should be called after all `Add` calls.**
- **A WaitGroup must not be copied after first use.**
- **Zero value is ready to use.**

## Common Mistakes

- Calling `Done` more or fewer times than `Add`.
- Calling `Wait` before all goroutines are started.
- Passing WaitGroup by value instead of by reference.

## Advanced Example: Waiting for HTTP Requests

```go
package main

import (
        "fmt"
        "net/http"
        "sync"
)

func fetch(url string, wg *sync.WaitGroup) {
        defer wg.Done()
        resp, err := http.Get(url)
        if err != nil {
                fmt.Println(err)
                return
        }
        fmt.Printf("%s: %s\n", url, resp.Status)
        resp.Body.Close()
}

func main() {
        var wg sync.WaitGroup
        urls := []string{"https://golang.org", "https://google.com", "https://github.com"}

        for _, url := range urls {
                wg.Add(1)
                go fetch(url, &wg)
        }

        wg.Wait()
        fmt.Println("All fetches complete")
}
```

## When to Use WaitGroup

- When you need to wait for multiple goroutines to finish.
- When you want to synchronize the completion of concurrent tasks.

## When Not to Use WaitGroup

- For signaling between goroutines (use channels instead).
- For waiting on a single event (use sync.Once or channels).

## Summary Table

| Method   | Description                              |
|----------|------------------------------------------|
| Add(n)   | Set number of goroutines to wait for     |
| Done()   | Signal that a goroutine is finished      |
| Wait()   | Block until all goroutines are done      |

## References

- [Go sync.WaitGroup Documentation](https://pkg.go.dev/sync#WaitGroup)
- [Go Concurrency Patterns](https://go.dev/doc/effective_go#concurrency)
