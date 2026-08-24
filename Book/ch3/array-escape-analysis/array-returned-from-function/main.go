package main

import "fmt"

func makeArray() [5]int {
	arr := [5]int{1, 2, 3, 4, 5}
	return arr
}

func main() {
	arr := makeArray()
	fmt.Println(arr)
}
