package main

import (
	"fmt"
	"time"
)

func main() {

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's a weekday")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("I am an int")
		case string:
			fmt.Println("I am a string")
		case bool:
			fmt.Println("I am a bool")
		default:
			fmt.Println("I am of unknown type", t)
		}
	}
	whoAmI(42)
	whoAmI("hello")
	whoAmI(true)
	whoAmI(3.14)
	whoAmI([]int{1, 2, 3})
	whoAmI(map[string]int{"one": 1, "two": 2})
}
