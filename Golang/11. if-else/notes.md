# Golang `if-else` Statement Notes

## Overview

The `if-else` statement in Go is used for conditional execution of code blocks. It allows you to execute certain code only if a specified condition is true, and optionally execute alternative code if the condition is false.

---

## Syntax

```go
if condition {
    // code to execute if condition is true
} else {
    // code to execute if condition is false
}
```

### Example

```go
age := 18
if age >= 18 {
    fmt.Println("You are an adult.")
} else {
    fmt.Println("You are a minor.")
}
```

---

## `if` Statement Only

You can use `if` without `else`:

```go
if x > 0 {
    fmt.Println("x is positive")
}
```

---

## `if-else if-else` Ladder

Multiple conditions can be checked using `else if`:

```go
score := 85
if score >= 90 {
    fmt.Println("Grade: A")
} else if score >= 80 {
    fmt.Println("Grade: B")
} else if score >= 70 {
    fmt.Println("Grade: C")
} else {
    fmt.Println("Grade: D or F")
}
```

---

## Short Statement in `if`

You can declare and initialize variables within the `if` statement:

```go
if n := len("hello"); n > 3 {
    fmt.Println("String is long enough")
}
```
- The variable `n` is only available inside the `if` and `else` blocks.

---

## No Parentheses

- Go does **not** require parentheses around the condition.
- Curly braces `{}` are **mandatory**.

---

## Nested `if-else`

You can nest `if-else` statements:

```go
num := 10
if num > 0 {
    if num%2 == 0 {
        fmt.Println("Positive even number")
    } else {
        fmt.Println("Positive odd number")
    }
} else {
    fmt.Println("Non-positive number")
}
```

---

## Comparison Operators

Common operators used in conditions:
- `==` (equal)
- `!=` (not equal)
- `>` (greater than)
- `<` (less than)
- `>=` (greater than or equal to)
- `<=` (less than or equal to)

---

## Logical Operators

Combine multiple conditions:
- `&&` (AND)
- `||` (OR)
- `!` (NOT)

Example:

```go
if age >= 18 && age <= 65 {
    fmt.Println("Working age")
}
```

---

## Best Practices

- Keep conditions simple and readable.
- Use short variable declarations in `if` when needed.
- Avoid deeply nested `if-else` blocks; consider using `switch` for multiple conditions.

---

## Summary Table

| Statement Type         | Usage Example                                 |
|------------------------|-----------------------------------------------|
| Simple if              | `if x > 0 { ... }`                            |
| if-else                | `if x > 0 { ... } else { ... }`               |
| if-else if-else ladder | `if ... else if ... else { ... }`             |
| Short statement        | `if v := f(); v > 0 { ... }`                  |
| Nested if              | `if ... { if ... { ... } }`                   |

---

## Complete Example

```go
package main

import "fmt"

func main() {
    if x := 7; x%2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }
}
```

---

## Conclusion

The `if-else` statement is fundamental for controlling program flow in Go. Mastering its syntax and usage is essential for writing effective Go programs.