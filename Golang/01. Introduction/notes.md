# Go Language Overview

## Compilation & Execution
- Compiled language - produces native machine code
- Go tool can run file directly without VM: `go run main.go`
- Executables are OS-specific:
    - Windows: `.exe`
    - Linux/Mac: binary file

## Use Cases
- System applications
- Web applications
- Cloud services
- Examples:
    - Docker (container platform)
    - Kubernetes (container orchestration)
    - CockroachDB (distributed SQL database)

## Language Features
### Missing Features (compared to other languages)
- No traditional try-catch blocks
    ```go
    // Error handling in Go
    if err != nil {
            return err
    }
    ```
- Heavy reliance on lexer for code formatting
    - Enforces strict formatting rules
    - Example: automatic semicolon insertion

### Notable Features
- Strong static typing
- Built-in concurrency support
- Garbage collection
- Fast compilation