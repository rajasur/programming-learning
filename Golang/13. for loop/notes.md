# Go For Loop - Complete Notes
---
## What is a For Loop?
A for loop is the only looping construct in Go. It's used to execute a block of code repeatedly based on a condition.
---
## Why Use For Loops?
- **Iteration**: Repeat code execution multiple times
- **Collection Processing**: Iterate over arrays, slices, maps, strings
- **Conditional Repetition**: Execute code while a condition is true
- **Counter-based Operations**: Perform operations a specific number of times
---
## When to Use For Loops?
- Processing collections of data
- Implementing algorithms that require repetition
- Reading files line by line
- Generating sequences
- Implementing retry mechanisms
---
## For Loop Syntax Variations

### 1. Traditional For Loop (C-style)
```go
for initialization; condition; post {
    // code block
}

// Example
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

### 2. While-style For Loop
```go
for condition {
    // code block
}

// Example
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

### 3. Infinite For Loop
```go
for {
    // code block
    // use break to exit
}

// Example
for {
    fmt.Println("Infinite loop")
    break // exits the loop
}
```

### 4. Range-based For Loop
```go
// For slices/arrays
for index, value := range collection {
    // code block
}

// For maps
for key, value := range mapVariable {
    // code block
}

// For strings (iterates over runes)
for index, char := range stringVariable {
    // code block
}
```
---
## Practical Examples

### Iterating Over Arrays/Slices
```go
numbers := []int{1, 2, 3, 4, 5}

// With index and value
for i, num := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", i, num)
}

// Only value
for _, num := range numbers {
    fmt.Println(num)
}

// Only index
for i := range numbers {
    fmt.Println("Index:", i)
}
```

### Iterating Over Maps
```go
person := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 35,
}

for name, age := range person {
    fmt.Printf("%s is %d years old\n", name, age)
}
```

### Iterating Over Strings
```go
text := "Hello, 世界"
for i, char := range text {
    fmt.Printf("Index: %d, Character: %c\n", i, char)
}
```
---
## Control Flow Statements

### Break Statement
```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // exits the loop
    }
    fmt.Println(i)
}
```

### Continue Statement
```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // skips current iteration
    }
    fmt.Println(i) // prints odd numbers only
}
```

### Labeled Break/Continue
```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer // breaks out of outer loop
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
```
---
## Common Patterns

### Nested Loops
```go
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        fmt.Printf("(%d,%d) ", i, j)
    }
    fmt.Println()
}
```

### Loop with Custom Step
```go
for i := 0; i < 20; i += 2 {
    fmt.Println(i) // prints even numbers
}
```

### Reverse Iteration
```go
numbers := []int{1, 2, 3, 4, 5}
for i := len(numbers) - 1; i >= 0; i-- {
    fmt.Println(numbers[i])
}
```
---
## Best Practices
- Use range for iterating over collections
- Use meaningful variable names for loop counters
- Avoid modifying loop variables inside the loop body
- Use break and continue judiciously
- Consider using labels for complex nested loops
- Prefer for-range over manual indexing when possible

## Common Pitfalls
- **Reference Issue**: Loop variables are reused
```go
// Wrong
var funcs []func()
for i := 0; i < 3; i++ {
    funcs = append(funcs, func() {
        fmt.Println(i) // will print 3, 3, 3
    })
}

// Correct
for i := 0; i < 3; i++ {
    i := i // create new variable
    funcs = append(funcs, func() {
        fmt.Println(i) // will print 0, 1, 2
    })
}
```

- **Infinite Loops**: Always ensure loop condition will eventually become false
- **Off-by-one Errors**: Check boundary conditions carefully