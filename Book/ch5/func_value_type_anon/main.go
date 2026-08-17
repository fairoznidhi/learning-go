package main

import "fmt"

// triple := func(x int) int {
// 	return x * 3
// } --->compile error: cannot use := outside function

// func()  {
// 	fmt.Println("Setup complete")
// }() --->compile error

var triple = func(x int)int{
	return x * 3
}

func main() {
	// Example of an anonymous function
	double := func(x int) int {
		return x * 2
	}
	fmt.Println(double(5)) 

	// Example of a function value
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(3, 4)) 
	fmt.Println(triple(5)) 
}