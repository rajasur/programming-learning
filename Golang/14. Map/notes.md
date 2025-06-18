# Maps in Go

## Introduction
- A map is a built-in data type in Go that associates keys with values
- Similar to dictionaries in Python or hash tables in other languages
- Maps are unordered collections of key-value pairs

## Syntax and Declaration
```go
// Declaration
var mapName map[KeyType]ValueType

// Initialization
mapName := make(map[KeyType]ValueType)

// Direct initialization with values
mapName := map[KeyType]ValueType{
    key1: value1,
    key2: value2,
}
```

## Key Features
1. **Dynamic Size**
   - Maps can grow and shrink at runtime
   - No fixed size like arrays

2. **Key Types**
   - Keys must be comparable types
   - Common key types: integers, strings, booleans
   - Structs and arrays can be keys if comparable
   - Slices and maps cannot be keys

3. **Value Types**
   - Can be any type including custom types
   - Can be another map or slice

## Operations
1. **Adding/Updating Elements**
   ```go
   mapName[key] = value
   ```

2. **Accessing Elements**
   ```go
   value := mapName[key]
   value, exists := mapName[key] // Check if key exists
   ```

3. **Deleting Elements**
   ```go
   delete(mapName, key)
   ```

4. **Length of Map**
   ```go
   length := len(mapName)
   ```

## Important Considerations
1. **Nil Maps**
   - Default value of map is nil
   - Cannot add elements to nil map
   - Must initialize before use

2. **Thread Safety**
   - Maps are not safe for concurrent access
   - Use sync.Map for concurrent operations

3. **Performance**
   - O(1) average time complexity for operations
   - Performance depends on hash function quality

4. **Memory Management**
   - Maps grow automatically
   - Memory is managed by garbage collector

## Best Practices
1. Use meaningful key types
2. Initialize maps with expected size when known
3. Check for key existence before operations
4. Clear maps by reassigning to empty map
5. Use range to iterate over maps

## Common Patterns
```go
// Iteration
for key, value := range mapName {
    // Process key and value
}

// Check key existence
if value, exists := mapName[key]; exists {
    // Key exists, use value
}

// Map of maps
map[string]map[string]int

// Map as set
map[string]bool
```

## Error Handling
- Accessing non-existent key returns zero value
- Modifying nil map causes panic
- Type assertions must be handled carefully with interface values

## Use Cases
1. Caching and memoization
2. Counting occurrences
3. Building lookup tables
4. Implementing sets
5. Graph representations