# Go Enums - Comprehensive Notes
---

## What are Enums?
Enums (enumerations) are a way to define a set of named constants that represent a fixed collection of values. Go doesn't have built-in enum support like other languages, but we can simulate them using constants and `iota`.

---

## When to Use Enums?
- When you have a fixed set of related constants
- To improve code readability and maintainability
- To prevent invalid values from being assigned
- For state machines, status codes, or configuration options
- When you need type safety for predefined values

---

## Why Use Enums?
- **Type Safety**: Prevents assignment of invalid values
- **Code Clarity**: Makes code more readable and self-documenting
- **Maintainability**: Easy to add/modify related constants
- **Performance**: Constants are resolved at compile time
- **IDE Support**: Better autocomplete and refactoring

---
## How to Implement Enums in Go

### Basic Enum with iota
```go
type Status int

const (
    Pending Status = iota
    Approved
    Rejected
    Cancelled
)
```

### String-based Enums
```go
type Color string

const (
    Red   Color = "red"
    Green Color = "green"
    Blue  Color = "blue"
)
```

### Enum with Custom Values
```go
type Priority int

const (
    Low    Priority = 1
    Medium Priority = 5
    High   Priority = 10
)
```

### Enum with String Method
```go
type Day int

const (
    Sunday Day = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func (d Day) String() string {
    days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
    if d < Sunday || d > Saturday {
        return "Unknown"
    }
    return days[d]
}
```

### Advanced Enum Patterns
```go
// Bitwise flags
type Permission int

const (
    Read Permission = 1 << iota
    Write
    Execute
)

// Enum validation
func (s Status) IsValid() bool {
    return s >= Pending && s <= Cancelled
}

// Enum from string
func StatusFromString(s string) (Status, error) {
    switch s {
    case "pending":
        return Pending, nil
    case "approved":
        return Approved, nil
    case "rejected":
        return Rejected, nil
    case "cancelled":
        return Cancelled, nil
    default:
        return 0, fmt.Errorf("invalid status: %s", s)
    }
}
```

## Best Practices
- Use typed constants instead of raw integers/strings
- Implement String() method for better debugging
- Add validation methods when needed
- Use meaningful names for enum types and values
- Group related constants together
- Consider using generate tools for complex enums