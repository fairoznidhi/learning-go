package main

import (
	"errors"
	"fmt"
)

func readConfig() error {
	baseErr := errors.New("file not found")
	return fmt.Errorf("readConfig failed: %w", baseErr)
}

func readConfigNoWrap() error {
	baseErr := errors.New("file not found")
	return fmt.Errorf("readConfig failed: %v", baseErr)
}

func main() {
	err := readConfig()
	if err != nil {
		fmt.Println(err)
		unwrappedErr := errors.Unwrap(err)
		if unwrappedErr != nil {
			fmt.Println(unwrappedErr)
		}
	}
	fmt.Println("-----")
	err2 := readConfigNoWrap()
	if err2 != nil {
		fmt.Println(err2)
		unwrappedErr := errors.Unwrap(err2)
		if unwrappedErr != nil {
			fmt.Println(unwrappedErr)
		}
	}
}
