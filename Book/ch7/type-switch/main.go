package main

import "fmt"

func describe(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("It's an integer %d\n", v)

	case string:
		fmt.Printf("It's a string %s\n", v)
	default:
		fmt.Println("Unknown type")
	}
}
func main() {
	describe(42)
	describe("Hello, Go!")
	describe(true)
}
