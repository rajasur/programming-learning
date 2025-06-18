# Arrays in Go Programming

## ✅ Introduction

An **array** is a fixed-size, homogeneous data structure that holds elements of the same type. Arrays in Go have a fixed length, which is part of their type.

```go
var arr [5]int // an array of 5 integers
```

## Default Values
When arrays are declared without initialization, elements get these default values:
- `int`: 0
- `float64`: 0.0
- `string`: "" (empty string)
- `bool`: false
- `pointer`: nil
- `struct`: zero value for each field

## Basic Syntax
```go
var arrayName [size]type
```

## Key Characteristics
1. Fixed Length
2. Zero-based indexing
3. Homogeneous elements
4. Value type (copying an array creates a new copy)
5. Default zero values for elements

## Declaration Examples
```go
// Method 1: Declaration with size (zero values)
var numbers [5]int  // [0,0,0,0,0]
var floats [3]float64 // [0.0,0.0,0.0]
var texts [3]string // ["","",""]
var flags [2]bool // [false,false]

// Method 2: Declaration with initialization
colors := [3]string{"Red", "Green", "Blue"}

// Method 3: Let compiler count elements
sizes := [...]int{1, 2, 3, 4, 5}
```

---

## 🔢 Array Declaration

### 1. **Basic Declaration**

```go
var a [3]int // default values will be 0
```

### 2. **With Initialization**

```go
var b = [3]int{1, 2, 3}
```

### 3. **Short Declaration**

```go
c := [3]int{4, 5, 6}
```

### 4. **Inferred Length**

```go
d := [...]int{7, 8, 9, 10} // Go infers length as 4
```

---

## 📍 Accessing Elements

Use the index to get or set an element.

```go
fmt.Println(a[0])  // get first element
a[1] = 100         // set second element
```

---

## 🔁 Iterating Over Arrays

### 1. **Using for loop**

```go
for i := 0; i < len(a); i++ {
    fmt.Println(a[i])
}
```

### 2. **Using for-range loop**

```go
for index, value := range b {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

---

## 📏 Array Length

Use `len()` to get the length of an array.

```go
fmt.Println(len(c)) // prints 3
```

---

## 🧬 Array is Value Type

In Go, arrays are **value types**. Assigning one array to another **copies all elements**.

```go
a := [3]int{1, 2, 3}
b := a
b[0] = 100

fmt.Println(a) // [1 2 3]
fmt.Println(b) // [100 2 3]
```

---

## 🧪 Multi-dimensional Arrays

You can declare arrays with more than one dimension.

```go
var matrix [2][3]int = [2][3]int{
    {1, 2, 3},
    {4, 5, 6},

nums := [2][2]int {
    {3,4},
    {5,6},
}
}
```

### Accessing Multi-dimensional Elements:

```go
fmt.Println(matrix[1][2]) // prints 6
```

---

## 🔄 Array of Structs

```go
type Employee struct {
    ID   int
    Name string
}

var team [2]Employee = [2]Employee{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
}
```

---

## 🧠 Copying Arrays

Arrays are copied **by value**, not by reference.

```go
x := [2]int{10, 20}
y := x
y[0] = 99

fmt.Println(x) // [10 20]
fmt.Println(y) // [99 20]
```

---

## 💡 Arrays vs Slices

| Feature     | Array          | Slice                        |
| ----------- | -------------- | ---------------------------- |
| Size        | Fixed          | Dynamic                      |
| Value Type  | Yes            | Reference Type               |
| Declaration | `[n]T`         | `[]T`                        |
| Copy        | Creates a copy | Points to same backing array |

---

## ❗ Common Mistakes

1. **Out of Range Index**

   ```go
   a := [3]int{1, 2, 3}
   fmt.Println(a[3]) // panic: index out of range
   ```

2. **Thinking Arrays are Reference Types**
   Arrays are **not like Java/C# arrays** – they are **copied**, not passed by reference.

---

## 🧮 Practical Example

```go
package main

import "fmt"

func main() {
    // Declare and initialize an array
    scores := [5]int{10, 20, 30, 40, 50}

    // Print elements using range
    for i, val := range scores {
        fmt.Printf("Index %d: Value %d\n", i, val)
    }

    // Total sum
    var sum int
    for _, val := range scores {
        sum += val
    }
    fmt.Println("Total:", sum)
}
```

---

## 📚 Summary

| Concept               | Example                                  |
| --------------------- | ---------------------------------------- |
| Declare               | `var a [5]int`                           |
| Initialize            | `a := [3]int{1,2,3}`                     |
| Access element        | `a[0] = 10`, `val := a[1]`               |
| Length                | `len(a)`                                 |
| Iteration             | `for i := 0; i < len(a); i++`            |
| Copy Behavior         | Arrays are copied, not referenced        |
| Multi-dimensional     | `[2][3]int`                              |
| Difference from Slice | Arrays are fixed-length, slices flexible |

---
