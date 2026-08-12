package main

import "fmt"
import "math/rand"

func main() {
	var randNums []int

	for i := 0; i < 100; i++ {
		randNum := rand.Intn(101) //[0,101)
		randNums = append(randNums, randNum)
	}
	
	for _, num := range randNums{
		switch {
		case num%2==0 && num%3==0:
			fmt.Println("Six!", num)
		case num%2==0:
			fmt.Println("Two!", num)
		case num%3==0:
			fmt.Println("Three!", num)
		default:
			fmt.Println("Never mind", num)
		}
	}
}
