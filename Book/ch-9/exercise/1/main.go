package main

import (
	"errors"
	"fmt"
)

var ErrInvalid error = errors.New("invalid ID")

func main() {
	derivedErr := fmt.Errorf("wrapped error %w", ErrInvalid)
	if errors.Is(derivedErr, ErrInvalid) {
		fmt.Println("Error matched")
	}

}
