package main

import (
	"errors"
	"fmt"
)

type StatusErr struct {
	Status  int
	Message string
	Err     error
}

func (se StatusErr) Error() string {
	return se.Message
}

func (se StatusErr) Unwrap() error {
	return se.Err
}

func readConfig() error {
	baseErr := errors.New("file not found")
	return StatusErr{
		Status:  404,
		Message: "readConfig failed",
		Err:     baseErr,
	}
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
}
