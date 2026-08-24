package main

import "fmt"

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func makePointer[T any](t T) *T {
	return &t
}

func doubleIt[T int | float64](t T) T {
	return t * 2
}

func main() {
	p := person{
		FirstName:  "Pat",
		MiddleName: makePointer("Perry"),
		LastName:   "Peterson",
	}
	fmt.Println(*p.MiddleName)

	result := doubleIt(21)
	fmt.Println(result)
}

// # command-line-arguments
// ./main.go:11:6: can inline makePointer[go.shape.string]
// ./main.go:15:6: can inline doubleIt[go.shape.int]
// ./main.go:15:6: can inline doubleIt[int]
// ./main.go:11:6: can inline makePointer[string]
// ./main.go:22:26: inlining call to makePointer[go.shape.string]
// ./main.go:25:13: inlining call to fmt.Println
// ./main.go:27:20: inlining call to doubleIt[go.shape.int]
// ./main.go:28:13: inlining call to fmt.Println
// ./main.go:15:6: inlining call to doubleIt[go.shape.int]
// ./main.go:11:6: inlining call to makePointer[go.shape.string]
// ./main.go:25:13: ... argument does not escape
// ./main.go:25:14: *p.MiddleName escapes to heap
// ./main.go:28:13: ... argument does not escape
// ./main.go:28:14: result escapes to heap
// ./main.go:11:25: moved to heap: t
// ./main.go:11:6: moved to heap: t
