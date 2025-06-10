// Golang program to calculate the sum of all array elements
package main

func sum(arr []int) int {
	cal := 0
	for _, value := range arr {
		cal += value
	}
	return cal
}
