package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	// ignoring the index value
	for _, v := range evenVals {
		fmt.Println(v)
	}

	fmt.Println("Map iteration")

	uniqueNames := map[string]bool{"Fred": true, "Wilma": true, "Raul": true}
	fmt.Println(uniqueNames) //--> ascending order by key

	for k := range uniqueNames {
		fmt.Println(k)
	}
}
