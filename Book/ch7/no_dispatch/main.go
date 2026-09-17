package main

import "fmt"

type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}
func (i Inner) hello() string {
	return "Hello from Inner"
}

func (i Inner) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outer struct {
	Inner
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}
func (o Outer) hello() string {
	return "Hello from Outer"
}

func main() {
	o := Outer{Inner: Inner{A: 5}, S: "Hello"}
	fmt.Println(o.Double())
	fmt.Println(o.hello())
	fmt.Println(o.Inner.hello())
}
