package main

import "fmt"

func sumArray() int {
	arr := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {
	total := sumArray()
	fmt.Println(total)
}

// # command-line-arguments
// ./main.go:5:6: can inline sumArray
// ./main.go:15:19: inlining call to sumArray
// ./main.go:16:13: inlining call to fmt.Println
// ./main.go:6:14: []int{...} does not escape
// ./main.go:15:19: []int{...} does not escape
// ./main.go:16:13: ... argument does not escape
// ./main.go:16:14: total escapes to heap. --->gets passed into fmt.Println's interface-based variadic parameter, has to be heap-allocated
