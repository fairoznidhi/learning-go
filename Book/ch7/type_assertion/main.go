package main

import "fmt"

type Shape interface {
	Area() float64
}
type Perimeter interface {
	Perimeter() float64
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func describeShape(s Shape) {
	fmt.Println(s.Area())
	perimeter, ok := s.(Perimeter)
	if ok {
		fmt.Println(perimeter.Perimeter())
	} else {
		fmt.Println("The shape does not implement the Perimeter interface.")
	}
}

func main() {
	s := Square{Side: 5}
	describeShape(s)
	c := Circle{Radius: 3}
	describeShape(c)
}
