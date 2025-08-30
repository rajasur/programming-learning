# Go Closures - Complete Notes
---
## What are Closures?

A closure is a function that captures and retains access to variables from its outer (enclosing) scope, even after the outer function has returned. In Go, closures are created when an inner function references variables from an outer function.

---
## When to Use Closures?

- **State preservation**: When you need to maintain state between function calls
- **Factory functions**: Creating specialized functions with pre-configured behavior
- **Event handlers**: Capturing context for callback functions
- **Iterators**: Maintaining iteration state
- **Decorators**: Wrapping functions with additional behavior
- **Configuration**: Creating functions with embedded configuration

---
## Why Use Closures?

- **Encapsulation**: Keep variables private and controlled
- **Code reusability**: Create specialized functions from generic templates
- **Stateful behavior**: Maintain state without global variables
- **Elegant solutions**: Cleaner code for certain patterns
- **Functional programming**: Enable higher-order function patterns

---
## Basic Closure Example

```go
package main

import "fmt"

func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    c := counter()
    fmt.Println(c()) // 1
    fmt.Println(c()) // 2
    fmt.Println(c()) // 3
}
```
---
## Multiple Closures with Independent State

```go
func main() {
    c1 := counter()
    c2 := counter()
    
    fmt.Println(c1()) // 1
    fmt.Println(c2()) // 1
    fmt.Println(c1()) // 2
    fmt.Println(c2()) // 2
}
```
---
## Closure with Parameters

```go
func multiplier(factor int) func(int) int {
    return func(n int) int {
        return n * factor
    }
}

func main() {
    double := multiplier(2)
    triple := multiplier(3)
    
    fmt.Println(double(5)) // 10
    fmt.Println(triple(5)) // 15
}
```
---
## Closures in Loops (Common Pitfall)

### Wrong Way:
```go
func main() {
    funcs := make([]func(), 3)
    for i := 0; i < 3; i++ {
        funcs[i] = func() {
            fmt.Println(i) // Will print 3, 3, 3
        }
    }
    
    for _, f := range funcs {
        f()
    }
}
```

### Correct Way:
```go
func main() {
    funcs := make([]func(), 3)
    for i := 0; i < 3; i++ {
        i := i // Create new variable
        funcs[i] = func() {
            fmt.Println(i) // Will print 0, 1, 2
        }
    }
    
    for _, f := range funcs {
        f()
    }
}
```
---
## Advanced Examples

### Event Handler with Context
```go
func createHandler(name string) func(string) {
    return func(event string) {
        fmt.Printf("%s handling %s\n", name, event)
    }
}
```

### Configuration Factory
```go
func createLogger(prefix string, debug bool) func(string) {
    return func(message string) {
        if debug {
            fmt.Printf("[DEBUG] %s: %s\n", prefix, message)
        } else {
            fmt.Printf("%s: %s\n", prefix, message)
        }
    }
}
```

### Accumulator Pattern
```go
func accumulator(initial int) func(int) int {
    sum := initial
    return func(n int) int {
        sum += n
        return sum
    }
}
```
---
## Memory Considerations

- Closures keep references to captured variables
- Variables remain in memory as long as closure exists
- Can lead to memory leaks if not handled properly
- Use pointers carefully in closures

---
## Best Practices

1. **Keep closures simple**: Avoid capturing too many variables
2. **Be mindful of loop variables**: Use local copies when needed
3. **Consider memory usage**: Don't capture large objects unnecessarily
4. **Use closures for state**: When you need controlled, private state
5. **Prefer closures over globals**: For maintaining state between calls

---
## Common Use Cases

- **Middleware functions**: HTTP middleware with configuration
- **Callback functions**: Event handlers with context
- **Factory patterns**: Creating specialized functions
- **State machines**: Maintaining state transitions
- **Decorators**: Adding behavior to existing functions
- **Currying**: Partial function application

---
## Performance Notes

- Closures have slight overhead compared to regular functions
- Variable capture involves heap allocation
- Garbage collection considers closure references
- Generally negligible impact for most applications