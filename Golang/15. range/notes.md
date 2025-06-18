# The `range` Keyword in Go

## Overview
The `range` keyword in Go is used to iterate over various data structures. It provides a simple and efficient way to loop through elements in slices, arrays, strings, maps, and channels.

## Syntax
```go
for index, value := range collection {
    // code block
}
```

## Use Cases

### 1. Slices and Arrays
```go
// Basic iteration
nums := []int{1, 2, 3, 4, 5}
for i, num := range nums {
    fmt.Printf("Index: %d, Value: %d\n", i, num)
}

// Using only index
for i := range nums {
    // work with i
}

// Using only value
for _, num := range nums {
    // work with num
}
```

### 2. Strings
```go
// Iterates over Unicode code points
str := "Hello, 世界"
for i, char := range str {
    fmt.Printf("position %d: %c\n", i, char)
}
```

### 3. Maps
```go
// Key-value iteration
myMap := map[string]int{"a": 1, "b": 2}
for key, value := range myMap {
    fmt.Printf("Key: %s, Value: %d\n", key, value)
}

// Keys only
for key := range myMap {
    // work with key
}
```

### 4. Channels
```go
// Reading from channel
ch := make(chan int)
for value := range ch {
    // process value until channel closes
}
```

## Important Notes
1. For arrays and slices:
   - Index is the integer index
   - Value is a copy of the element

2. For strings:
   - Index is the byte index
   - Value is the Unicode code point (rune)

3. For maps:
   - Iteration order is not guaranteed
   - Key and value are copies

4. For channels:
   - Range continues until channel is closed
   - Only receives values, no index

## Best Practices
1. Use blank identifier (_) when you don't need index or key
2. Be aware that range creates copies of values
3. For large structs, use pointer receivers to avoid copying
4. Consider performance implications when ranging over large data structures

## Common Pitfalls
1. Modifying slice while ranging
2. Accessing map elements while ranging
3. Not closing channels when using range
4. Assuming fixed iteration order for maps

## Performance Considerations
- Range creates copies of elements
- For large structs, use pointers
- Consider using traditional for loops for simple numeric iterations