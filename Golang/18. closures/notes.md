# Closures in Go

## What is a Closure?
A **closure** is a function value that references variables from outside its body. The function may access and modify the variables even after the outer function has finished executing.

## Why Use Closures?
- Encapsulate state
- Create function factories
- Implement callbacks and handlers

## How Closures Work in Go
When you define an inner function that refers to variables from an outer function, Go creates a closure. The referenced variables are captured by the closure and remain accessible.

## Example: Basic Closure

```go
package main

import "fmt"

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    add := adder()
    fmt.Println(add(10)) // 10
    fmt.Println(add(5))  // 15
}
```
- `sum` is captured by the returned function.
- Each call to `add` updates and remembers `sum`.

## Multiple Closures with Independent State

```go
func main() {
    a := adder()
    b := adder()
    fmt.Println(a(1)) // 1
    fmt.Println(a(2)) // 3
    fmt.Println(b(1)) // 1 (separate state)
}
```

## Closures and Goroutines

Be careful when using closures in goroutines, especially in loops:

```go
for i := 0; i < 3; i++ {
    go func() {
        fmt.Println(i) // May print 3, 3, 3
    }()
}
```
**Fix:** Pass variable as parameter:

```go
for i := 0; i < 3; i++ {
    go func(n int) {
        fmt.Println(n)
    }(i)
}
```

## Closures as Function Arguments

Closures can be passed as arguments:

```go
func apply(f func(int) int, x int) int {
    return f(x)
}

func main() {
    double := func(n int) int { return n * 2 }
    fmt.Println(apply(double, 5)) // 10
}
```

## Closures and Variable Capture

Closures capture variables **by reference**, not by value. This means the variable's value can change after the closure is created.

## Summary

- Closures are functions that capture variables from their environment.
- Useful for maintaining state, callbacks, and function factories.
- Be careful with variable capture in loops and goroutines.
