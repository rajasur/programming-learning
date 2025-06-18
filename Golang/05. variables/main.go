package main

import (
	"fmt"
)

// Package level variable - can be accessed throughout the package
var Raja string = "value"

// Package level constant - immutable value
// Public constant (starts with capital letter, accessible from other packages)
const LoginToken string = "adhhofohnfeiuwfh"

// Multiple constants can be declared together
const (
	Pi       = 3.14159
	MaxUsers = 1000
	AppName  = "MyGoApp"
)

func main() {
	// Basic variable declaration with type
	var username string = "Raja"
	var isLoggedIn bool = true

	//inferred type declaration
	var paerson_age = 30 // Type is inferred as int
	var isAdult = true
	// Printing the values
	fmt.Println("Person Age is:", paerson_age)
	fmt.Println("Is Adult:", isAdult)

	// Printing variables and their types
	fmt.Println("Username:", username)
	fmt.Printf("Variable type: %T\n", username)
	fmt.Println("Login status:", isLoggedIn)
	fmt.Printf("Variable type: %T\n", isLoggedIn)

	// Integer variable
	var smallValue int = 23456
	fmt.Println("Integer value:", smallValue)
	fmt.Printf("Type: %T\n", smallValue)

	// Short declaration operator (:=)
	// Type is inferred automatically
	smallFloat := 2345.118765
	fmt.Printf("Float value: %f\n", smallFloat)
	fmt.Printf("Type: %T\n", smallFloat)

	// Default values (zero values) in Go
	var (
		defaultInt    int
		defaultFloat  float64
		defaultString string
		defaultBool   bool
	)
	fmt.Println("\nDefault values:")
	fmt.Printf("Integer: %v\n", defaultInt)   // 0
	fmt.Printf("Float: %v\n", defaultFloat)   // 0.0
	fmt.Printf("String: %q\n", defaultString) // ""
	fmt.Printf("Boolean: %v\n", defaultBool)  // false

	// Implicit type declaration
	var website = "rajasur.in" // Type is inferred as string
	fmt.Println("\nWebsite:", website)

	// Short declaration with multiple variables
	name, age := "John", 25
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Constants cannot be changed
	fmt.Println("\nConstants:")
	fmt.Println("Login Token:", LoginToken)
	fmt.Println("Pi:", Pi)
	fmt.Println("Max Users:", MaxUsers)
	fmt.Println("App Name:", AppName)

	// Type conversion example
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println("\nType conversion:")
	fmt.Printf("int %v converted to float64 %v and uint %v\n", i, f, u)
}
