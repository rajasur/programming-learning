# Struct Embedding in Go
---

## What is Struct Embedding?

Struct embedding is a feature in Go that allows you to include one struct type inside another struct type without explicitly naming the field. The embedded struct's fields and methods become directly accessible through the embedding struct, providing a form of inheritance-like behavior.

```go
type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person  // Embedded struct
    ID      int
    Salary  float64
}
```
---
## Why Use Struct Embedding?

- **Code Reusability**: Avoid duplicating common fields across multiple structs
- **Composition over Inheritance**: Go's approach to achieving inheritance-like behavior
- **Method Promotion**: Embedded struct methods are automatically available on the embedding struct
- **Cleaner Code**: Reduces boilerplate and creates more maintainable code
- **Interface Satisfaction**: Embedding can help structs satisfy interfaces
---
## When to Use Struct Embedding?

- When you have common fields/behaviors across multiple types
- Building hierarchical relationships (Employee is a Person)
- Extending existing types without modification
- Creating wrapper types that need to satisfy specific interfaces
- Implementing the decorator pattern

## Basic Syntax

```go
type Address struct {
    Street   string
    City     string
    ZipCode  string
}

type Person struct {
    Name    string
    Age     int
    Address // Embedded struct
}
```
---
## Accessing Embedded Fields

```go
p := Person{
    Name: "John",
    Age:  30,
    Address: Address{
        Street:  "123 Main St",
        City:    "New York",
        ZipCode: "10001",
    },
}

// Direct access (promoted fields)
fmt.Println(p.Street)   // "123 Main St"
fmt.Println(p.City)     // "New York"

// Explicit access
fmt.Println(p.Address.Street) // "123 Main St"
```
---
## Method Promotion

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return a.Name + " makes a sound"
}

type Dog struct {
    Animal
    Breed string
}

func main() {
    d := Dog{
        Animal: Animal{Name: "Buddy"},
        Breed:  "Golden Retriever",
    }
    
    fmt.Println(d.Speak()) // "Buddy makes a sound"
}
```
---
## Multiple Embedding

```go
type Walker struct {
    Speed int
}

func (w Walker) Walk() string {
    return "Walking at speed: " + strconv.Itoa(w.Speed)
}

type Swimmer struct {
    Depth int
}

func (s Swimmer) Swim() string {
    return "Swimming at depth: " + strconv.Itoa(s.Depth)
}

type Amphibian struct {
    Walker
    Swimmer
    Name string
}
```
---
## Name Conflicts and Resolution

```go
type A struct {
    Value int
}

type B struct {
    Value string
}

type C struct {
    A
    B
    // Value field is ambiguous - must be accessed explicitly
}

func main() {
    c := C{
        A: A{Value: 42},
        B: B{Value: "hello"},
    }
    
    // fmt.Println(c.Value) // Error: ambiguous selector
    fmt.Println(c.A.Value) // 42
    fmt.Println(c.B.Value) // "hello"
}
```
---
## Interface Satisfaction Through Embedding

```go
type Writer interface {
    Write([]byte) (int, error)
}

type Logger struct {
    *os.File // Embedding pointer to File
    prefix   string
}

func (l *Logger) Log(message string) {
    l.Write([]byte(l.prefix + message))
}

// Logger satisfies Writer interface through embedding
```
---
## Embedding Interfaces

```go
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type ReadWriter interface {
    Reader
    Writer
}
```
---
## Best Practices

1. **Use embedding for "is-a" relationships**
2. **Prefer composition over deep embedding hierarchies**
3. **Be careful with name conflicts**
4. **Use pointer embedding when the embedded type is large**
5. **Document embedded behavior clearly**
---
## Common Patterns

### Extending Third-party Types

```go
type MyWriter struct {
    *bytes.Buffer
    writeCount int
}

func (mw *MyWriter) Write(p []byte) (n int, err error) {
    mw.writeCount++
    return mw.Buffer.Write(p)
}
```

### Mixin Pattern

```go
type Timestamped struct {
    CreatedAt time.Time
    UpdatedAt time.Time
}

type User struct {
    Timestamped
    Name  string
    Email string
}

type Product struct {
    Timestamped
    Name  string
    Price float64
}
```
---
## Limitations

- No true inheritance (no polymorphism)
- Potential name conflicts with multiple embedding
- Can make code harder to understand if overused
- Circular embedding is not allowed
- Cannot embed the same type multiple times directly
---
## Example: Complete Implementation

```go
package main

import (
    "fmt"
    "time"
)

type Entity struct {
    ID        int
    CreatedAt time.Time
}

func (e Entity) GetID() int {
    return e.ID
}

type Person struct {
    Entity
    Name  string
    Email string
}

func (p Person) String() string {
    return fmt.Sprintf("Person{ID: %d, Name: %s, Email: %s}", 
        p.ID, p.Name, p.Email)
}

type Employee struct {
    Person
    Department string
    Salary     float64
}

func main() {
    emp := Employee{
        Person: Person{
            Entity: Entity{ID: 1, CreatedAt: time.Now()},
            Name:   "Alice",
            Email:  "alice@example.com",
        },
        Department: "Engineering",
        Salary:     75000,
    }
    
    fmt.Println(emp.String())        // Uses Person's String method
    fmt.Println(emp.GetID())         // Uses Entity's GetID method
    fmt.Println(emp.Name)            // Direct access to embedded field
    fmt.Println(emp.Entity.ID)       // Explicit access
}
```