### Six Main Points about GO Programming Language

---

1. **Statically Typed Language**

   ```go
   var x int = 10
   fmt.Println(x)
   ```

   In Go, the type of variable `x` is known at compile time.

2. **Strong Typed Language**

   ```go
   var x int = 10
   var y float64 = 20.0
   // This will cause a compile-time error
   // z := x + y
   ```

   Go does not allow implicit type conversion, ensuring type safety.

3. **Go is Compiled**

   ```sh
   go build main.go
   ./main
   ```

   Go code is compiled into a binary executable.

4. **Fast Compile Time**

   ```sh
   time go build main.go
   ```

   Go is designed to have a fast compilation process.

5. **Built-In Concurrency**

   ```go
   go func() {
        fmt.Println("Hello from a goroutine")
   }()
   ```

   Go has built-in support for concurrent programming using goroutines.

6. **Simplicity**

   ```go
   package main

   import "fmt"

   func main() {
        fmt.Println("Hello, World!")
   }
   ```

   Go emphasizes simplicity and clarity in its syntax and design.

### How to Intiate your Go program, build it and run it?

---

To initiate your Go program, follow these steps:

1. **Create a New Go File**

   Create a new file with a `.go` extension, for example, `main.go`.

   ```sh
   touch main.go
   ```

2. **Write Your Go Code**

   Open the `main.go` file in your preferred text editor and write your Go code. For example:

   ```go
   package main

   import "fmt"

   func main() {
        fmt.Println("Hello, World!")
   }
   ```

3. **Initialize a New Go Module**

   Initialize a new Go module in your project directory. This will create a `go.mod` file.

   ```sh
   go mod init your-module-name
   ```

4. **Build Your Go Program**

   Use the `go build` command to compile your Go code into an executable binary.

   ```sh
   go build main.go
   ```

   This will create an executable file named `main` (or `main.exe` on Windows).

5. **Run Your Go Program**

   Execute the compiled binary to run your Go program.

   ```sh
   ./main
   ```

   You should see the output:

   ```
   Hello, World!
   ```

By following these steps, you can easily initiate, build, and run your Go program.

### Running Your Go Program Without Building

---

You can also run your Go program without explicitly building it using the `go run` command. This command compiles and runs your Go code in one step.

```sh
go run main.go
```

This will compile `main.go` and immediately execute the resulting binary. You should see the output:

```
Hello, World!
```

Using `go run` is convenient for quick iterations and testing during development.

### Go Laxer

---

Go Laxer is a tool that helps in analyzing and visualizing Go code. It provides insights into code structure, dependencies, and potential issues. Here are some key features:

1. **Code Analysis**

   Go Laxer performs static code analysis to identify potential issues and code smells.

2. **Dependency Visualization**

   It visualizes the dependencies between different packages and modules in your Go project.

3. **Code Metrics**

   Go Laxer provides various code metrics such as cyclomatic complexity, lines of code, and more.

4. **Integration with CI/CD**

   It can be integrated into your CI/CD pipeline to ensure code quality and maintainability.

5. **User-Friendly Interface**

   Go Laxer offers a user-friendly interface to navigate and explore your Go codebase.

By using Go Laxer, developers can maintain high code quality and gain better insights into their Go projects.
**Note:** Go laxer also remove terminated semicolon( ; ) automatically.

### Golangs Types

---

- Case insensitive; almost
- Variable type should be known in advance.
- Everything is a Type

#### Types

- String
- Bool
- Integer
  - uint8
  - uint64
  - int8
  - int64
  - uintptr
- Floating
  - float32
  - float64
- Complex
- Array
- Slices
- Maps
- Structs
- Pointers
- Functions
- Channels
