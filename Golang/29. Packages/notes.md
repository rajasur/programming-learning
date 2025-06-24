# Go Packages

## What is a Package?

A **package** in Go is a way to organize and reuse code. Every Go source file belongs to a package. Packages help manage dependencies and prevent naming conflicts.

## Package Declaration

At the top of every Go file, declare the package name:

```go
package main
```

- `main` is a special package for executable programs.
- Other packages are for libraries or reusable code.

## Importing Packages

Use the `import` keyword to use code from other packages:

```go
import "fmt"
```

You can import multiple packages:

```go
import (
    "fmt"
    "math"
)
```

## Creating Your Own Package

1. **Create a folder** with your package name.
2. **Add Go files** with `package <name>` at the top.
3. **Exported identifiers** (functions, types, variables) must start with a capital letter.

**Example:**

**mathutils/add.go**
```go
package mathutils

func Add(a, b int) int {
    return a + b
}
```

**main.go**
```go
package main

import (
    "fmt"
    "yourmodule/mathutils"
)

func main() {
    sum := mathutils.Add(2, 3)
    fmt.Println(sum) // Output: 5
}
```

## The `main` Package and Function

- The `main` package is the entry point for executables.
- Must have a `func main()`.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## Standard Library Packages

Go provides many built-in packages, e.g.:
- `fmt` for formatted I/O
- `math` for mathematics
- `os` for operating system functions
- `net/http` for HTTP servers

**Example:**

```go
import "math"

func main() {
    fmt.Println(math.Sqrt(16)) // Output: 4
}
```

## Aliasing Imports

You can rename imports:

```go
import m "math"

func main() {
    fmt.Println(m.Pi)
}
```

## Dot Import

Imports all exported identifiers directly (not recommended):

```go
import . "fmt"

func main() {
    Println("Hello") // No fmt. prefix needed
}
```

## Blank Identifier Import

Import for side-effects only:

```go
import _ "net/http/pprof"
```

## Package Initialization

- Each package can have an `init()` function.
- Runs before `main()` or before the package is used.

```go
func init() {
    fmt.Println("Initializing package")
}
```

## Summary

- Packages organize code and manage dependencies.
- Use `import` to use other packages.
- Exported names start with uppercase.
- The `main` package is for executables.
- Use the standard library for common tasks.

