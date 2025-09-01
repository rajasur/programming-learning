# Go Packages and Scope

## What are Packages?

A package in Go is a collection of source files in the same directory that are compiled together. Packages are Go's way of organizing and reusing code.

### Key Characteristics:
- Every Go source file belongs to a package
- Package name is declared at the top of each file
- All files in a directory must belong to the same package
- Package `main` is special - it defines an executable program

## When to Use Packages?

### Use Cases:
- **Code Organization**: Group related functionality together
- **Code Reusability**: Share code across multiple programs
- **Namespace Management**: Avoid naming conflicts
- **Encapsulation**: Control access to functions and variables
- **Modular Development**: Break large programs into manageable pieces

## Why Use Packages?

### Benefits:
- **Maintainability**: Easier to maintain and update code
- **Testability**: Each package can be tested independently
- **Collaboration**: Teams can work on different packages simultaneously
- **Performance**: Only import what you need
- **Security**: Control visibility of internal implementation

## Package Scope Rules

### Exported vs Unexported Identifiers

```go
package mypackage

// Exported (public) - starts with uppercase
var PublicVariable = "accessible from other packages"
func PublicFunction() {}
type PublicStruct struct {}

// Unexported (private) - starts with lowercase
var privateVariable = "only accessible within this package"
func privateFunction() {}
type privateStruct struct {}
```

### Visibility Levels:
1. **Package-level**: Accessible within the same package
2. **Exported**: Accessible from other packages (capitalized)
3. **Unexported**: Only accessible within the same package (lowercase)

## Package Declaration and Import

### Package Declaration:
```go
package packagename
```

### Import Statements:
```go
import "fmt"                    // Standard library
import "github.com/user/repo"   // External package
import "./localpackage"         // Local package

// Multiple imports
import (
    "fmt"
    "os"
    "strings"
)

// Import with alias
import f "fmt"

// Blank import (for side effects)
import _ "database/sql/driver"
```

## Package Structure Example

```
project/
├── main.go
├── utils/
│   ├── math.go
│   └── string.go
└── models/
    └── user.go
```

### main.go:
```go
package main

import (
    "fmt"
    "./utils"
    "./models"
)

func main() {
    result := utils.Add(5, 3)
    user := models.NewUser("John")
    fmt.Println(result, user.Name)
}
```

### utils/math.go:
```go
package utils

// Exported function
func Add(a, b int) int {
    return a + b
}

// Unexported function
func multiply(a, b int) int {
    return a * b
}
```

## Best Practices

1. **Use descriptive package names**: Short, clear, lowercase
2. **Avoid package name conflicts**: Don't use standard library names
3. **Group related functionality**: Keep logically related code together
4. **Minimize package dependencies**: Avoid circular imports
5. **Use interfaces**: Define contracts between packages
6. **Document exported functions**: Use comments for public APIs

## Package Initialization

```go
package mypackage

import "fmt"

// Package-level variables
var packageVar = initializeVar()

// init function runs before main
func init() {
    fmt.Println("Package initialized")
}

func initializeVar() string {
    return "initialized"
}
```

## Common Pitfalls

- **Circular imports**: Package A imports B, B imports A
- **Incorrect capitalization**: Forgetting to capitalize exported identifiers
- **Package naming**: Using underscores or mixed case in package names
- **Large packages**: Putting too much functionality in one package