# Golang Functions - Detailed Notes

## What is a Function?
A function is a reusable block of code that performs a specific task. Functions help organize code, avoid repetition, and improve readability.

## Syntax

```go
func functionName(parameterList) returnType {
    // function body
}
```

- `func`: Keyword to declare a function.
- `functionName`: Name of the function.
- `parameterList`: List of parameters (optional).
- `returnType`: Type of value returned (optional).

## Example

```go
func add(a int, b int) int {
    return a + b
}
```

## Calling a Function

```go
result := add(2, 3) // result is 5
```

## Parameters

- **Single Parameter**: `func greet(name string)`
- **Multiple Parameters**: `func add(a int, b int)`
- **Same Type Parameters**: `func add(a, b int)`
- **Variadic Parameters**: Accepts zero or more arguments.

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

## Return Values

- **Single Return**: `func square(x int) int`
- **Multiple Returns**: `func divide(a, b int) (int, error)`

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

## Named Return Values

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```

## Anonymous Functions

Functions without a name, often used as closures.

```go
add := func(a, b int) int {
    return a + b
}
fmt.Println(add(3, 4))
```

## Functions as Values

Functions can be assigned to variables, passed as arguments, or returned from other functions.

```go
func operate(a, b int, op func(int, int) int) int {
    return op(a, b)
}
```

## Defer

`defer` schedules a function call to run after the function completes.

```go
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
// Output: hello
//         world
```

## Recursion

A function calling itself.

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```

## Methods

Functions with a receiver argument, associated with types.

```go
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}
```

## Summary

- Functions are fundamental in Go for code organization.
- Support for multiple parameters, multiple return values, named returns, variadic functions, anonymous functions, and methods.
- Functions can be used as first-class citizens.
