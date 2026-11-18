// Command inventory prints a generated UUID and a formatted, doubled number.
package main

import (
	"fmt"

	"example.com/inventory/price"

	"github.com/google/uuid"
	format "github.com/learning-go-book-2e/package_example/do-format"
)

func main() {
	// Print a new random UUID.
	fmt.Println(uuid.New())
	// Double 5 using the local price package.
	num := price.Double(5)
	// Format the number using the third-party do-format package.
	output := format.Number(num)
	fmt.Println(output)
}
