package main

import (
	"errors"
	"fmt"
)

type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}

const ErrPkgA = Sentinel("not found")
const ErrPkgB = Sentinel("not found")

func main() {
	errA := errors.New("not found")
	errB := errors.New("not found")
	errC := errA
	if errA == errB {
		fmt.Println("errA and errB are equal")
	}
	if errA == errC {
		fmt.Println("errA and errC are equal")
	}
	if errB == errC {
		fmt.Println("errB and errC are equal")
	}

	if ErrPkgA == ErrPkgB {
		fmt.Println("Errors are equal")
	} else {
		fmt.Println("Errors are not equal")
	}
}
