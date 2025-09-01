# Go (Golang) Packages & Commands Notes

------------------------------------------------------------------------

## 📌 Go Packages

### What is a Package?

-   A **package** in Go is a collection of related Go source files that
    are compiled together.
-   Each Go program is made up of **packages**.
-   The main package (`package main`) defines the entry point.

### When to Use?

-   Use packages to **organize code** into reusable units.
-   Use the `main` package when creating executables.
-   Use standard library or third-party packages when you want pre-built
    functionality.

### Why?

-   Promotes **code reusability, modularity, and maintainability**.
-   Avoids duplication by separating logic.

------------------------------------------------------------------------

## 📦 Types of Packages

### 1. **Standard Library Packages**

-   Predefined by Go, examples: `fmt`, `os`, `net/http`, `math`, `time`,
    etc.
-   **When?** → Use when common functionality is needed (printing, file
    I/O, networking, etc.).\
-   **Why?** → Saves time and ensures optimized, well-tested code.

### 2. **Custom Packages**

-   User-defined, reusable code stored in modules.
-   **When?** → Use for your own reusable components.\
-   **Why?** → Encapsulation and modularity.

### 3. **Third-Party Packages**

-   Installed via `go get <module>`.
-   **When?** → Use when functionality isn't available in stdlib.\
-   **Why?** → Faster development with community tools.

------------------------------------------------------------------------

## 📌 Go Commands

### 🔹 1. `go run`

-   **What?** → Compiles and runs Go programs without creating a
    binary.\
-   **When?** → For quick testing and debugging.\
-   **Why?** → Faster iteration during development.

``` bash
go run main.go
```

------------------------------------------------------------------------

### 🔹 2. `go build`

-   **What?** → Compiles source code into an executable binary.\
-   **When?** → To produce executables for deployment.\
-   **Why?** → Provides optimized binary for distribution.

``` bash
go build main.go
```

------------------------------------------------------------------------

### 🔹 3. `go install`

-   **What?** → Compiles and installs the binary into `$GOPATH/bin` or
    module's bin path.\
-   **When?** → To install tools and binaries globally.\
-   **Why?** → Easier access to executables from terminal.

``` bash
go install github.com/some/tool@latest
```

------------------------------------------------------------------------

### 🔹 4. `go fmt`

-   **What?** → Automatically formats code according to Go standards.\
-   **When?** → Before committing code.\
-   **Why?** → Ensures consistency and readability.

``` bash
go fmt ./...
```

------------------------------------------------------------------------

### 🔹 5. `go vet`

-   **What?** → Reports suspicious constructs (static analysis).\
-   **When?** → During development and CI/CD pipelines.\
-   **Why?** → Catches bugs early (e.g., Printf mismatches).

``` bash
go vet ./...
```

------------------------------------------------------------------------

### 🔹 6. `go test`

-   **What?** → Runs unit tests (`*_test.go`).\
-   **When?** → To verify correctness of code.\
-   **Why?** → Encourages test-driven development.

``` bash
go test ./...
```

------------------------------------------------------------------------

### 🔹 7. `go mod`

-   **What?** → Manages Go modules (dependencies).\
-   **When?** → For projects using modules (`go.mod`).\
-   **Why?** → Handles versioning and dependency management.

``` bash
go mod init myproject
go mod tidy
go mod download
```

------------------------------------------------------------------------

### 🔹 8. `go get`

-   **What?** → Downloads and installs remote packages.\
-   **When?** → When adding external dependencies.\
-   **Why?** → Integrates third-party modules.

``` bash
go get github.com/gin-gonic/gin
```

------------------------------------------------------------------------

### 🔹 9. `go clean`

-   **What?** → Removes object files and cached binaries.\
-   **When?** → To free space or rebuild cleanly.\
-   **Why?** → Prevents conflicts from old builds.

``` bash
go clean
```

------------------------------------------------------------------------

### 🔹 10. `go doc`

-   **What?** → Displays documentation for a package or symbol.\
-   **When?** → While learning new packages.\
-   **Why?** → Easy access to reference docs.

``` bash
go doc fmt.Println
```

------------------------------------------------------------------------

### 🔹 11. `go list`

-   **What?** → Lists available packages and modules.\
-   **When?** → To inspect dependencies.\
-   **Why?** → Useful for dependency analysis.

``` bash
go list all
```

------------------------------------------------------------------------

### 🔹 12. `go env`

-   **What?** → Shows Go environment variables.\
-   **When?** → For debugging setup issues.\
-   **Why?** → Helps configure Go environment.

``` bash
go env
```

------------------------------------------------------------------------

### 🔹 13. `go version`

-   **What?** → Displays installed Go version.\
-   **When?** → To check Go installation.\
-   **Why?** → Ensures compatibility.

``` bash
go version
```

------------------------------------------------------------------------

### 🔹 14. `go generate`

-   **What?** → Runs code generation commands marked with
    `//go:generate`.\
-   **When?** → For scaffolding, mocks, etc.\
-   **Why?** → Automates repetitive code tasks.

``` bash
go generate ./...
```

------------------------------------------------------------------------

### 🔹 15. `go tool`

-   **What?** → Runs various Go tools (cover, trace, pprof).\
-   **When?** → For debugging and performance analysis.\
-   **Why?** → Advanced diagnostics.

``` bash
go tool cover -html=coverage.out
```

------------------------------------------------------------------------

## 📌 Go Workspace & Modules

-   **Workspace** → Collection of Go projects managed with modules.\
-   **Modules** → Unit of versioned Go code, defined in `go.mod`.\
-   **When?** → Always use modules in Go ≥ 1.16.\
-   **Why?** → Standardized dependency management.

------------------------------------------------------------------------

## 📌 Summary (What, When, Why)

| Command         | What it Does         | When to Use         | Why Important           |
|-----------------|---------------------|---------------------|-------------------------|
| `go run`        | Runs program without binary | Quick testing        | Fast feedback           |
| `go build`      | Builds binary        | Deployments         | Optimized executables   |
| `go install`    | Installs binary globally | CLI tools            | Easy execution          |
| `go fmt`        | Formats code         | Before commits      | Code consistency        |
| `go vet`        | Static analysis      | Development & CI    | Catch bugs early        |
| `go test`       | Runs tests           | Development         | Ensures correctness     |
| `go mod`        | Manages dependencies | Projects            | Dependency mgmt         |
| `go get`        | Downloads packages   | Add dependencies    | Reuse code              |
| `go clean`      | Removes cache        | Clean rebuilds      | Avoid conflicts         |
| `go doc`        | Shows docs           | Learning            | Quick reference         |
| `go list`       | Lists packages       | Dependency checks   | Transparency            |
| `go env`        | Shows env vars       | Debug setup         | Fix issues              |
| `go version`    | Go version           | Compatibility       | Validation              |
| `go generate`   | Code generation      | Scaffolding         | Automation              |
| `go tool`       | Runs tools           | Debugging           | Profiling & analysis    |

------------------------------------------------------------------------
