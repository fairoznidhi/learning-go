package main

import "fmt"

func main() {
	true := 5
	fmt.Println(true)
	// true is in universal block, so it can be shadowed
	// but it cannot be used to compare a boolean value

	// predefined identifiers: --> can be shadowed --> live in universal block
	// types:
	// bool, byte, complex64, complex128, error, float32, float64, int, int8, int16, int32, int64, rune, string, uint, uint8, uint16,
	// uint32, uint64, uintptr

	// constants:
	// true, false, iota

	// functions:
	// append, cap, close, complex, copy, delete, imag, len, make, new, panic, print, println, real, recover

	// keywords: --> defined in Go spec's lexical grammar --> cannot be shadowed
	// break, case, chan, const, continue, default, defer, else, fallthrough, for, func, go, goto, if, import, interface,
	// map, package, range, return, select, struct, switch, type, var

}
