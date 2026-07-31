package main

import "fmt"

func main() {
	fmt.Println("\nLiterals in Go")
	fmt.Println("Integer literal:", 42)

	fmt.Println("Floating-point literal:", 3.14)

	fmt.Println("Rune literal----->:", 'A')

	fmt.Println("String literal:", "Hello, Go!")
	// interpreted string literal -> as they interpret rune literals(numeric and escape)
	fmt.Println("Tab:\there")
	fmt.Println("Newline:\nthere")
	fmt.Println("Backslash:\\there")
	fmt.Println("Double quote: \"there\"")
	fmt.Println("Single quote: 'there'")

	fmt.Println("Raw literal(no escape)")
	fmt.Println(`Greetings and
"Salutations"`)

}
