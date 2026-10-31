package main

import (
	"fmt"

	"github.com/fairoznidhi/workspace_lib"
)

func main() {
	result := workspace_lib.AddNums(3, 4)
	fmt.Println("The sum is:", result)

	result2 := workspace_lib.SubNums(10, 5)
	fmt.Println("The difference is:", result2)
}
