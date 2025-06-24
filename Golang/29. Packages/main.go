package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rajasur/programming-learning/GO/auth"
	"github.com/rajasur/programming-learning/GO/user"
)

func main() {
	auth.LoginWithCredentials("RajaSur", "sap@123456")

	session := auth.GetSession()
	println("Session:", session)
	user := user.User{
		Email: "user@email.com",
		Name:  "John Doe",
	}
	fmt.Println(user.Email, user.Name)
	color.Green(user.Email)
	color.Red(user.Name)
}
