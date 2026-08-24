package main

import "fmt"

func makePointer() *[5]int {
	arr := [5]int{1, 2, 3, 4, 5}
	return &arr
}
func main() {
	arr := makePointer()
	fmt.Println(*arr)
}

// # command-line-arguments
// ./main.go:5:6: can inline makePointer
// ./main.go:10:20: inlining call to makePointer
// ./main.go:11:13: inlining call to fmt.Println
// ./main.go:6:2: moved to heap: arr
// ./main.go:11:13: ... argument does not escape
// ./main.go:11:14: *arr escapes to heap
