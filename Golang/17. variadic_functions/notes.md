# Variadic Functions in Go
---
## What are Variadic Functions?

Variadic functions are functions that can accept a variable number of arguments of the same type. In Go, they are defined using three dots (`...`) before the parameter type.

---
## When to Use Variadic Functions

- When you need to accept an unknown number of arguments
- Creating utility functions like `fmt.Printf`, `fmt.Println`
- Building functions that work with collections of data
- When you want to make function calls more flexible and user-friendly
---
## Why Use Variadic Functions

- **Flexibility**: Allows functions to handle varying numbers of inputs
- **Convenience**: Eliminates need for slice creation in function calls
- **Cleaner API**: Makes function calls more intuitive
- **Backward Compatibility**: Easy to extend existing functions
---
## Syntax

```go
func functionName(params ...type) returnType {
    // function body
}
```
---
## Basic Example

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Usage
result1 := sum(1, 2, 3)        // 6
result2 := sum(1, 2, 3, 4, 5)  // 15
result3 := sum()               // 0
```
---
## Passing Slices to Variadic Functions

```go
numbers := []int{1, 2, 3, 4, 5}
result := sum(numbers...)  // Spread operator
```
---
## Mixed Parameters

```go
func greet(greeting string, names ...string) {
    for _, name := range names {
        fmt.Printf("%s, %s!\n", greeting, name)
    }
}

greet("Hello", "Alice", "Bob", "Charlie")
```
---
## Key Points

- Variadic parameter must be the last parameter
- Inside function, variadic parameter is treated as a slice
- Can pass zero or more arguments
- Use `...` to pass slice elements as individual arguments
- Empty variadic parameter creates empty slice, not nil