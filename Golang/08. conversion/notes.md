**complete, detailed guide to Golang (Go) type conversion**, covering **all key conversion topics** with **examples, best practices**.

---

## 📘 Go Language: Type Conversion (Type Casting)

Go is **statically typed**, which means variables have a defined type. Type **conversion** (also called **type casting**) is required when converting between different types, like from `int` to `float64`, `string` to `int`, etc.

Go does **not allow implicit conversions** — all conversions must be **explicit**.

---

## ✅ Table of Contents

1. [Basic Syntax](#1-basic-syntax)
2. [Integer to Float and Vice Versa](#2-integer-to-float-and-vice-versa)
3. [String to Int/Float and Vice Versa](#3-string-to-intfloat-and-vice-versa)
4. [Rune and Byte Conversions](#4-rune-and-byte-conversions)
5. [Interface Type Assertions](#5-interface-type-assertions)
6. [Custom Type Conversions](#6-custom-type-conversions)
7. [Best Practices](#7-best-practices)
8. [Summary Table](#8-summary-table)

---

## 1. 🔤 Basic Syntax

```go
T(v)
```

Where `T` is the type you want to convert `v` into.

### Example:

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

---

## 2. 🔢 Integer ↔ Float

### Integer to Float

```go
var i int = 10
var f float64 = float64(i)
fmt.Println(f)  // Output: 10.0
```

### Float to Integer (decimal truncated)

```go
var f float64 = 9.87
var i int = int(f)
fmt.Println(i)  // Output: 9
```

---

## 3. 🔁 String ↔ Int/Float

You must use the `strconv` package.

### ✅ Import:

```go
import "strconv"
```

---

### String to Integer

```go
s := "123"
i, err := strconv.Atoi(s)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(i)  // Output: 123
```

---

### Integer to String

```go
i := 456
s := strconv.Itoa(i)
fmt.Println(s)  // Output: "456"
```

---

### String to Float

```go
s := "3.1415"
f, err := strconv.ParseFloat(s, 64)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(f)  // Output: 3.1415
```

---

### Float to String

```go
f := 2.718
s := strconv.FormatFloat(f, 'f', 6, 64)
fmt.Println(s)  // Output: "2.718000"
```

---

### String to Bool

```go
b, err := strconv.ParseBool("true")
fmt.Println(b) // Output: true
```

---

### Bool to String

```go
b := true
s := strconv.FormatBool(b)
fmt.Println(s) // Output: "true"
```

---

## 4. 🧩 Rune and Byte Conversions

### String to Byte Slice

```go
s := "Go"
b := []byte(s)
fmt.Println(b)  // Output: [71 111]
```

### Byte Slice to String

```go
b := []byte{71, 111}
s := string(b)
fmt.Println(s)  // Output: "Go"
```

---

### Character to Rune

```go
var ch byte = 'A'
r := rune(ch)
fmt.Println(r)  // Output: 65
```

---

## 5. 🧠 Interface{} and Type Assertion

Used when you have values of `interface{}` type and want to convert them back to their actual type.

### Example:

```go
var val interface{} = "Hello"

str, ok := val.(string)
if ok {
    fmt.Println("Value:", str)
} else {
    fmt.Println("Not a string")
}
```

---

## 6. 🛠️ Custom Type Conversions

You can convert between **custom types** if they share the same base type.

### Example:

```go
type Celsius float64
type Fahrenheit float64

var c Celsius = 100
f := Fahrenheit(c)  // Explicit conversion
fmt.Println(f)      // Output: 100
```

---

## 7. 🧼 Best Practices

✅ Always check for errors when converting strings
✅ Avoid unnecessary conversions (especially in tight loops)
✅ Prefer `strconv` for reliable string-number conversions
✅ Understand truncation when converting from float to int
✅ Use type assertions carefully with `interface{}`

---

## 8. 📋 Summary Table

| From → To        | Conversion Method                    | Example                            |
| ---------------- | ------------------------------------ | ---------------------------------- |
| int → float      | `float64(i)`                         | `float64(10)`                      |
| float → int      | `int(f)`                             | `int(3.14)` (truncates to 3)       |
| string → int     | `strconv.Atoi(s)`                    | `"123"` → `123`                    |
| int → string     | `strconv.Itoa(i)`                    | `456` → `"456"`                    |
| string → float   | `strconv.ParseFloat(s, 64)`          | `"2.71"` → `2.71`                  |
| float → string   | `strconv.FormatFloat(f, 'f', 6, 64)` | `2.71` → `"2.710000"`              |
| string → bool    | `strconv.ParseBool("true")`          | `"true"` → `true`                  |
| bool → string    | `strconv.FormatBool(true)`           | `true` → `"true"`                  |
| string ↔ byte\[] | `[]byte(s)` / `string([]byte)`       | `"Go"` ↔ `[71 111]`                |
| interface{}      | `val.(type)`                         | `val.(string)` with type assertion |

---

## 🧪 Example: All Conversions in One Program

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    i := 42
    f := float64(i)
    s := strconv.Itoa(i)

    str := "99"
    i2, _ := strconv.Atoi(str)
    f2, _ := strconv.ParseFloat("3.14", 64)

    b := []byte("GoLang")
    str2 := string(b)

    fmt.Println("int → float:", f)
    fmt.Println("int → string:", s)
    fmt.Println("string → int:", i2)
    fmt.Println("string → float:", f2)
    fmt.Println("string → []byte:", b)
    fmt.Println("[]byte → string:", str2)
}
```

---

Would you like this Markdown saved to a file (`conversion.md`) or want a downloadable PDF?
