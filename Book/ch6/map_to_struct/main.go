//	book := map[string]interface{}{
//			"title":  "The Go Programming Language",
//			"author": "Donovan and Kernighan",
//			"pages":  400,
//		}
package main

import "fmt"

type Book struct {
	Title  string //capitalize-exported //lowercase-unexported
	Author string
	Pages  int
}

func main() {
	// book := Book{"The Go Programming Language", "Donovan and Kernighan", 400}
	book := Book{Title: "The Go Programming Language", Author: "Donovan and Kernighan", Pages: 400}
	fmt.Println(book)
}
