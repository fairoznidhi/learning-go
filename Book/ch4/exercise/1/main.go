package main

import "fmt"
import "math/rand"

func main() {
	var randNums []int

	for i := 0; i < 100; i++ {
		randNum := rand.Intn(101) //[0,101)
		randNums = append(randNums, randNum)
	}
	fmt.Println(randNums)
}
