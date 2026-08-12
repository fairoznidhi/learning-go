package main

import "fmt"

func main() {
	var total int
	for i:=0; i<10;i++{
		total:= total + i // total is shadowed everytime in the loop, so it will always be 0 + i
	}
}