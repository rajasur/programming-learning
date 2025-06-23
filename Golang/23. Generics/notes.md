# Golang Generics Notes

## Introduction
Generics in Go allow you to write flexible and reusable functions, types, and data structures that can work with any data type. Introduced in Go 1.18, generics use type parameters to achieve type safety without code duplication.

---

## Why Use Generics?
- **Code Reusability:** Write functions and types that work with any type.
- **Type Safety:** Catch errors at compile time.
- **Performance:** Avoid interface{} and type assertions.

---

## Syntax Overview

### Type Parameters
Type parameters are specified in square brackets `[]` after the function or type name.

```go
func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}
```
- `T` is a type parameter.
- `any` is a built-in constraint (alias for `interface{}`).

---

## Constraints

Constraints restrict the set of types that can be used as type parameters.

### Built-in Constraints
- `any`: Any type.
- `comparable`: Types that support `==` and `!=`.

### Custom Constraints
Define an interface with required methods or type sets.

```go
type Adder[T any] interface {
    Add(a, b T) T
}
```

Or using type sets (union):

```go
type Number interface {
    int | float64
}
```

---

## Generic Functions

```go
func Swap[T any](a, b T) (T, T) {
    return b, a
}
```

**Usage:**
```go
x, y := Swap[int](1, 2)
```

---

## Generic Types

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}
```

---

## Type Inference

Go can often infer type parameters:

```go
Swap(1, 2) // T inferred as int
```

---

## Multiple Type Parameters

```go
func Pair[A, B any](a A, b B) (A, B) {
    return a, b
}
```

---

## Methods with Type Parameters

```go
func (s *Stack[T]) Pop() T {
    n := len(s.items)
    item := s.items[n-1]
    s.items = s.items[:n-1]
    return item
}
```

---

## Generic Interfaces

```go
type Equaler[T any] interface {
    Equal(a, b T) bool
}
```

---

## Type Sets

Type sets define which types satisfy a constraint.

```go
type Number interface {
    ~int | ~float64
}
```
- `~` allows underlying types.

---

## Example: Generic Map Function

```go
func Map[T any, U any](s []T, f func(T) U) []U {
    result := make([]U, len(s))
    for i, v := range s {
        result[i] = f(v)
    }
    return result
}
```

---

## Limitations

- No operator overloading.
- No generic constants.
- Type parameters cannot be used for struct fields' names.

---

## Best Practices

- Use generics for truly generic code.
- Avoid overusing generics for simple cases.
- Name type parameters with single uppercase letters (`T`, `K`, `V`).

---

## References

- [Go Generics Proposal](https://go.dev/design/43651-type-parameters)
- [Go Blog: Generics](https://go.dev/blog/intro-generics)
