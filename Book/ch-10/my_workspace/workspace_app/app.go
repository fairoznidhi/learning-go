package main

import (
	"fmt"

	"github.com/fairoznidhi/workspace_lib"
)

func main() {
	var a, b int = 3, 4
	a = 9
	result := workspace_lib.AddNums(a, b)
	fmt.Println("The sum is:", result)

	result2 := workspace_lib.SubNums(10, 5)
	fmt.Println("The difference is:", result2)
}
