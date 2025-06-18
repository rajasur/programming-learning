# Slices in Go

## Introduction
A slice is a dynamic, flexible view into an array. Unlike arrays, slices can grow and shrink. They are one of the most commonly used data structures in Go.

## Syntax and Creation

### Basic Syntax
```go
var slice []datatype
```

### Different Ways to Create Slices

1. Using make function:
```go
slice := make([]int, length, capacity)
slice := make([]int, length) // capacity equals length
```

2. Using array:
```go
array := [5]int{1, 2, 3, 4, 5}
slice := array[1:4] // creates slice from index 1 to 3
```

3. Slice literal:
```go
slice := []int{1, 2, 3, 4, 5}
```

## Key Concepts

### Length and Capacity
- Length: Number of elements in the slice
- Capacity: Number of elements in the underlying array from the start of the slice
```go
slice := make([]int, 3, 5)
fmt.Println(len(slice)) // 3
fmt.Println(cap(slice)) // 5
```

### Slicing Operations
```go
slice := []int{1, 2, 3, 4, 5}
slice1 := slice[1:4]    // [2, 3, 4]
slice2 := slice[:3]     // [1, 2, 3]
slice3 := slice[2:]     // [3, 4, 5]
```

### Appending Elements
```go
slice := []int{1, 2, 3}
slice = append(slice, 4)        // [1, 2, 3, 4]
slice = append(slice, 5, 6, 7)  // [1, 2, 3, 4, 5, 6, 7]
```

### Copying Slices
```go
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)
```

## Important Behaviors

### Nil Slices
```go
var slice []int // nil slice
fmt.Println(slice == nil) // true
```

### Empty Slices
```go
slice := []int{} // empty slice
fmt.Println(slice == nil) // false
```

### Slice Memory Management
- Slices hold references to underlying array
- Memory leaks can occur if large arrays are referenced by small slices
```go
// Example of potential memory leak
array := [1000000]int{}
slice := array[0:1] // Still holds reference to large array
```

## Best Practices

1. Pre-allocate when size is known:
```go
slice := make([]int, 0, expectedSize)
```

2. Use copy() for slice duplication:
```go
newSlice := make([]int, len(oldSlice))
copy(newSlice, oldSlice)
```

3. Be careful with large slices of slices:
```go
// Instead of
slice = append(slice, element)
// Consider pre-allocating
slice = make([]int, 0, expectedSize)
```

## Common Operations

### Removing Elements
```go
// Remove from middle
slice = append(slice[:i], slice[i+1:]...)

// Remove from start
slice = slice[1:]

// Remove from end
slice = slice[:len(slice)-1]
```

### Extending Slices
```go
slice = append(slice, []int{4, 5, 6}...)
```

### Clear Slice
```go
slice = slice[:0] // Maintains capacity
```

## Performance Considerations

1. Avoid frequent resizing:
    - Pre-allocate with expected capacity
    - Use append wisely

2. Memory efficiency:
    - Use copy() to create new slices when needed
    - Be aware of underlying array references

3. Slice operations complexity:
    - Append: O(1) amortized
    - Copy: O(n)
    - Slicing: O(1)

## Common Pitfalls

1. Modifying through different slices:
```go
original := []int{1, 2, 3}
slice1 := original[:]
slice2 := original[:]
slice1[0] = 10 // Also affects original and slice2
```

2. Unexpected capacity growth:
```go
slice := make([]int, 0, 5)
slice = append(slice, 1, 2, 3, 4, 5)
slice = append(slice, 6) // Capacity doubles
```

3. Slice bounds out of range:
```go
slice := []int{1, 2, 3}
fmt.Println(slice[5]) // Runtime panic
```

## Built-in copy Function

The `copy` function is a built-in function that copies elements from a source slice to a destination slice:

```go
func copy(dst, src []Type) int
```

### Example with copy
```go
source := []int{1, 2, 3, 4}
dest := make([]int, 2)
copy(dest, source)
fmt.Println(dest)    // [1, 2]
fmt.Println(source)  // 2 (number of elements copied)
```

Notes:
- Returns number of elements copied
- Copies minimum of len(dst) or len(src) elements
- Safe for overlapping slices