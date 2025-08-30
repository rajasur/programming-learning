# Go Generics Notes

## What are Generics?

Generics allow you to write code that works with multiple types while maintaining type safety. They enable you to create functions, types, and methods that can operate on different data types without sacrificing compile-time type checking.

```go
// Without generics - type-specific functions
func PrintInt(s []int) {
    for _, v := range s {
        fmt.Println(v)
    }
}

func PrintString(s []string) {
    for _, v := range s {
        fmt.Println(v)
    }
}

// With generics - single function for multiple types
func Print[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}
```

## Why Use Generics?

1. **Code Reusability**: Write once, use with multiple types
2. **Type Safety**: Compile-time type checking prevents runtime errors
3. **Performance**: No boxing/unboxing or interface{} overhead
4. **Reduced Code Duplication**: Eliminate repetitive type-specific code
5. **Better API Design**: More expressive and flexible interfaces

## When to Use Generics?

- **Data Structures**: Slices, maps, trees, linked lists
- **Algorithms**: Sorting, searching, filtering functions
- **Utility Functions**: Functions that work with multiple types
- **Collections**: When you need type-safe containers
- **Mathematical Operations**: Functions that work with numeric types

### When NOT to Use Generics
- When working with a single, specific type
- Simple functions that don't benefit from type abstraction
- When interface{} is sufficient and performance isn't critical

## Generic Function Syntax

```go
// Basic generic function
func Swap[T any](a, b T) (T, T) {
    return b, a
}

// Multiple type parameters
func Compare[T, U comparable](a T, b U) bool {
    return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

// Using the functions
x, y := Swap[int](1, 2)
str1, str2 := Swap("hello", "world") // Type inference
```

## Type Constraints

### Built-in Constraints
```go
// any - equivalent to interface{}
func Process[T any](value T) T {
    return value
}

// comparable - types that support == and !=
func Equal[T comparable](a, b T) bool {
    return a == b
}
```

### Custom Constraints
```go
// Interface-based constraints
type Numeric interface {
    int | int8 | int16 | int32 | int64 |
    uint | uint8 | uint16 | uint32 | uint64 |
    float32 | float64
}

func Add[T Numeric](a, b T) T {
    return a + b
}

// Method-based constraints
type Stringer interface {
    String() string
}

func ToString[T Stringer](items []T) []string {
    result := make([]string, len(items))
    for i, item := range items {
        result[i] = item.String()
    }
    return result
}
```

## Generic Types

### Generic Structs
```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

// Usage
intStack := Stack[int]{}
intStack.Push(1)
intStack.Push(2)
```

### Generic Interfaces
```go
type Container[T any] interface {
    Add(T)
    Get() T
    Size() int
}

type List[T any] struct {
    items []T
}

func (l *List[T]) Add(item T) {
    l.items = append(l.items, item)
}

func (l *List[T]) Get() T {
    if len(l.items) == 0 {
        var zero T
        return zero
    }
    return l.items[0]
}

func (l *List[T]) Size() int {
    return len(l.items)
}
```

## Type Inference

Go can often infer types automatically:

```go
func Max[T comparable](a, b T) T {
    if a > b { // This won't compile - comparable doesn't support >
        return a
    }
    return b
}

// Better version with proper constraint
type Ordered interface {
    int | int8 | int16 | int32 | int64 |
    uint | uint8 | uint16 | uint32 | uint64 |
    float32 | float64 | string
}

func Max[T Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Usage with type inference
result := Max(10, 20)      // T inferred as int
maxStr := Max("a", "b")    // T inferred as string
```

## Practical Examples

### Generic Map Function
```go
func Map[T, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

// Usage
numbers := []int{1, 2, 3, 4}
strings := Map(numbers, func(n int) string {
    return fmt.Sprintf("Number: %d", n)
})
```

### Generic Filter Function
```go
func Filter[T any](slice []T, predicate func(T) bool) []T {
    var result []T
    for _, v := range slice {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}

// Usage
numbers := []int{1, 2, 3, 4, 5, 6}
evens := Filter(numbers, func(n int) bool {
    return n%2 == 0
})
```

### Generic Optional/Maybe Type
```go
type Optional[T any] struct {
    value *T
}

func Some[T any](value T) Optional[T] {
    return Optional[T]{value: &value}
}

func None[T any]() Optional[T] {
    return Optional[T]{value: nil}
}

func (o Optional[T]) IsSome() bool {
    return o.value != nil
}

func (o Optional[T]) Unwrap() T {
    if o.value == nil {
        panic("called Unwrap on None value")
    }
    return *o.value
}

func (o Optional[T]) UnwrapOr(defaultValue T) T {
    if o.value == nil {
        return defaultValue
    }
    return *o.value
}
```

## Best Practices

1. **Use meaningful constraint names**
2. **Prefer composition over complex constraints**
3. **Keep generic functions simple and focused**
4. **Use type inference when possible**
5. **Document generic types and their constraints**
6. **Test with multiple type instantiations**

## Common Pitfalls

1. **Zero values**: Always handle zero values properly
2. **Type constraints**: Ensure constraints match your operations
3. **Performance**: Generics can impact compile time
4. **Readability**: Don't over-generalize simple code

## golang.org/x/exp/constraints Package

```go
import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

func Sum[T constraints.Integer | constraints.Float](nums []T) T {
    var sum T
    for _, num := range nums {
        sum += num
    }
    return sum
}
```