# Golang `switch` Statement Notes

## Overview
The `switch` statement in Go is a control structure for multi-way branching. It evaluates an expression and executes the case that matches its value.

---

## Basic Syntax

```go
switch expression {
case value1:
    // code
case value2:
    // code
default:
    // code
}
```

---

## How It Works

- The `expression` is evaluated once.
- Each `case` is compared to the result.
- The first matching `case` block runs.
- If no cases match, the `default` block runs (if present).

---

## Example

```go
num := 2
switch num {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
default:
    fmt.Println("Other")
}
// Output: Two
```

---

## Switch Without Expression

- Omitting the expression makes it act like a series of `if-else` statements.

```go
switch {
case num < 0:
    fmt.Println("Negative")
case num == 0:
    fmt.Println("Zero")
default:
    fmt.Println("Positive")
}
```

---

## Multiple Expressions in a Case

- Separate values with commas.

```go
switch day := "Saturday"; day {
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Weekday")
}
```

---

## Fallthrough

- Use `fallthrough` to continue to the next case.

```go
num := 1
switch num {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two")
}
// Output:
// One
// Two
```

---

## Type Switch

- Used to compare types, especially with interfaces.

```go
var i interface{} = 10
switch v := i.(type) {
case int:
    fmt.Println("int:", v)
case string:
    fmt.Println("string:", v)
default:
    fmt.Println("unknown")
}
```

---

## Initialization Statement

- You can declare variables in the switch statement.

```go
switch x := getValue(); x {
case 1:
    // ...
}
```

---

## Notes

- Cases do **not** automatically fall through (unlike C/C++).
- Only one case runs unless `fallthrough` is used.
- Cases can use expressions (e.g., `case x > 10:` in expressionless switch).
- `default` is optional.
- Cases must be unique and constant values (except in type switches).

---

## Summary Table

| Feature                | Supported? | Example/Note                        |
|------------------------|------------|-------------------------------------|
| Expression             | Yes        | `switch x { ... }`                  |
| No expression          | Yes        | `switch { ... }`                    |
| Multiple case values   | Yes        | `case 1, 2:`                        |
| Fallthrough            | Yes        | `fallthrough` keyword               |
| Type switch            | Yes        | `switch v := i.(type) { ... }`      |
| Initialization         | Yes        | `switch x := f(); x { ... }`        |
| Default case           | Optional   | `default:`                          |

---

## References

- [Go Tour: Switch](https://tour.golang.org/flowcontrol/9)
- [Go Blog: Switch](https://blog.golang.org/switch)
