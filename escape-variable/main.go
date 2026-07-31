package main

import "fmt"

func main() {
	x := 5
	ptr := multiply(x)
	fmt.Println(*ptr)
}

func multiply(x int) *int {
	z := x * 100
	return &z
}
