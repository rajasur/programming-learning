# 📘 Go Language: User Input (Complete Guide)

In Go, user input can be received in multiple ways:
- Using `fmt.Scan`, `fmt.Scanf`, `fmt.Scanln` for terminal input
- Using `bufio.NewReader` for line-by-line or character input
- Using `os.Stdin` for reading raw bytes
- Using `flag` package for command-line arguments

## ✅ 1. Using `fmt` Package

### 📌 `fmt.Scan()`
Reads space-separated inputs.

```go
package main

import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("Enter your name and age: ")
    fmt.Scan(&name, &age)
    fmt.Printf("Name: %s\nAge: %d\n", name, age)
}
```

### 📌 `fmt.Scanln()`
Reads input until newline, but still space-separated.

```go
package main

import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("Enter your name and age: ")
    fmt.Scanln(&name, &age)
    fmt.Printf("Name: %s\nAge: %d\n", name, age)
}
```

### 📌 `fmt.Scanf()`
Reads formatted input using format specifiers.

```go
package main

import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("Enter your name and age (format: Name Age): ")
    fmt.Scanf("%s %d", &name, &age)
    fmt.Printf("Name: %s\nAge: %d\n", name, age)
}
```

## ✅ 2. Using `bufio` + `os.Stdin`

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your full name: ")
    name, _ := reader.ReadString('\n')
    fmt.Printf("Your name is: %s", name)
}
```

### 📌 String to Number Conversion

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your age: ")
    input, _ := reader.ReadString('\n')
    age, err := strconv.Atoi(strings.TrimSpace(input))
    if err != nil {
        fmt.Println("Invalid number!")
        return
    }
    fmt.Printf("You are %d years old.\n", age)
}
```

## ✅ 3. Using `flag` Package

```go
package main

import (
    "flag"
    "fmt"
)

func main() {
    name := flag.String("name", "Guest", "Enter your name")
    age := flag.Int("age", 0, "Enter your age")
    flag.Parse()
    fmt.Printf("Name: %s\nAge: %d\n", *name, *age)
}
```
```bash
go run main.go -name=Raja -age=25
```

## 🧪 Input Methods Summary

| Method | Use Case | Spaces | Line | Format |
|--------|----------|--------|------|--------|
| `fmt.Scan` | Simple values | No | No | No |
| `fmt.Scanln` | Line input | No | Yes | No |
| `fmt.Scanf` | Formatted | No | No | Yes |
| `bufio` | Full lines | Yes | Yes | No |
| `flag` | CLI args | Yes | Yes | Yes |

## 🛠️ Best Practices
1. Use `bufio.NewReader` for multiline input
2. Use `Scan` family for simple inputs
3. Handle conversion errors
4. Use `flag` for CLI tools
