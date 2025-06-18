package main

import (
	"fmt"
	"maps"
)

// maps -> hash, obj, dict
func main() {
	// Creating map
	m := make(map[string]string)

	// setting an element
	m["name"] = "Golang"
	m["area"] = "backend"

	// get an element
	fmt.Println("Name:", m["name"])
	fmt.Println("Area:", m["area"])
	fmt.Println(m["phone"]) // IMP: if key does not exist, it returns zero value of the type

	fmt.Println((len(m))) // Length of map is 2

	delete(m, "area")                      // delete an element
	fmt.Println("After deleting area:", m) // After deleting area: map[name:Golang]

	clear(m)                              // clear the map
	fmt.Println("After clearing map:", m) // After clearing map: map[]

	m1 := map[string]int{"price": 100, "quantity": 2}
	fmt.Println("Map m1:", m1) // Map m1: map[price:100 quantity:2]

	value, ok := m1["price"]              // check if key exists
	fmt.Println("Value of price:", value) // Value of price: 100
	if ok {
		fmt.Println("Price exists in m1:", m1["price"]) // Price exists in m1: 100
	} else {
		fmt.Println("Price does not exist in m1")
	}

	// maps package
	m2 := map[string]int{"price": 100, "quantity": 2}
	m3 := map[string]int{"price": 100, "quantity": 2}
	fmt.Println("m2 == m3:", maps.Equal(m3, m2)) // m2 == m3: true

}
