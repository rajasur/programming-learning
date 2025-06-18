# Golang Pointers Notes

## What is a Pointer?
- A pointer is a variable that stores the memory address of another variable.
- In Go, pointers are strongly typed.

## Declaring Pointers
```go
var p *int // p is a pointer to an int
```

## Zero Value of Pointers
- The zero value of a pointer is `nil`.

## Getting the Address of a Variable
- Use the `&` operator to get the address.
```go
a := 10
p := &a // p now holds the address of a
```

## Dereferencing a Pointer
- Use the `*` operator to access the value stored at the address.
```go
fmt.Println(*p) // prints 10
*p = 20         // changes a to 20
```

## Creating Pointers with `new`
- The `new` function allocates memory and returns a pointer.
```go
p := new(int) // p points to a newly allocated int, initialized to 0
```

## Pointers and Functions
- Pointers allow functions to modify variables outside their scope.
```go
func increment(n *int) {
    *n++
}
```

## Pointers to Structs
- Use pointers to structs to avoid copying and to modify struct fields.
```go
type Person struct { Name string }
p := &Person{Name: "Alice"}
p.Name = "Bob" // shorthand for (*p).Name
```

## No Pointer Arithmetic
- Go does **not** support pointer arithmetic (unlike C/C++).

## Nil Pointers
- Always check for `nil` before dereferencing a pointer to avoid runtime panics.

## Arrays vs Slices and Pointers
- Arrays are values; passing them copies the array.
- Slices are references; passing them does not copy the underlying array.

## Pointers to Pointers
- Go supports pointers to pointers, but their use is rare.

## When to Use Pointers
- To modify a variable inside a function.
- To avoid copying large structs or arrays.
- To share data between parts of a program.

## Example
```go
func swap(a, b *int) {
    *a, *b = *b, *a
}
```

## Best Practices
- Use pointers when necessary, but avoid excessive use.
- Prefer slices and maps for sharing data.
- Always initialize pointers before use.
