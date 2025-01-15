package main

import (
	"fmt"
)

var Raja string = "value"
// jwtToken := "value" // this will not work outside of function
const LoginToken string = "adhhofohnfeiuwfhfh" // It is kind of Public variable if the first letter is capital then it is public


func main() {
	fmt.Println("variables")
	var username string = "Raja"
	var isLoggedIn bool = true
	fmt.Println("Username: ", username)
	fmt.Printf("Variable is of type: %T\n", username)
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	var smallValue int = 23456
	fmt.Println(smallValue)
	fmt.Printf("Type: %T \n", smallValue)

	//walrus operator i.e :=
	smallFloat := 2345.118765
	print(smallFloat)
	fmt.Printf("Type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int 
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of Type: %T \n", anotherVariable)

	// implicit type
	var website = "rajasur.in"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}
