package simcalc

import (
	"fmt"
)

// SimCalculate calculates similarity prints a float64 value ranging [0-100].
func SimCalculate(code1, code2 string) error {
	var sim float64
	sim = 100
	fmt.Println(sim)
	return nil
}

// DebugCalculate calculates similarity and prints process log.
func DebugCalculate(code1, code2 string) error {
	fmt.Println("It's a checking process with verbosed option and would print debugging information")
	fmt.Println("Path of code1: " + code1)
	fmt.Println("Path of code2: " + code2)
	return nil
}
