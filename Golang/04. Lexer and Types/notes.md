# Lexer and Types in Go

## Lexical Elements

### 1. Comments
- Single-line comments: Start with `//`
- Multi-line comments: Enclosed between `/*` and `*/`

### 2. Tokens
- Identifiers
- Keywords
- Operators and Punctuation
- Literals

### 3. Keywords
Important keywords in Go:
```go
break    default     func    interface    select
case     defer       go      map          struct
chan     else        goto    package      switch
const    fallthrough if      range        type
continue for         import  return       var
```

## Types in Go

### 1. Basic Types
- **Numeric Types**
    - Integers: `int8`, `int16`, `int32`, `int64`, `int`
    - Unsigned: `uint8`, `uint16`, `uint32`, `uint64`, `uint`
    - Float: `float32`, `float64`
    - Complex: `complex64`, `complex128`
- **String Type**: `string`
- **Boolean Type**: `bool`

### 2. Composite Types
- **Array**: Fixed-length sequence
- **Slice**: Dynamic-length sequence
- **Map**: Key-value pairs
- **Struct**: User-defined type
- **Interface**: Method set

### 3. Type Declarations
```go
// Type alias
type MyInt = int

// Type definition
type Age int
```

### 4. Type Conversions
```go
var i int = 42
var f float64 = float64(i)
```

## Type Safety
- Go is statically typed
- Type checking happens at compile time
- Implicit type conversion is not allowed

## Zero Values
Each type has a default zero value:
- Numeric types: `0`
- Boolean: `false`
- String: `""`
- Pointer: `nil`
- Interface: `nil`
- Reference types: `nil`

## Type Inference
Go can infer types when declaring variables:
```go
x := 42        // int
y := "hello"   // string
z := true      // bool
```