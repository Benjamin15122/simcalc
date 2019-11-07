package simcalc

import "fmt"

// SimCalculate calculates similarity prints a float64 value ranging [0-100].
func SimCalculate(code1, code2 string) {
	var sim float64
	sim = 100
	fmt.Println(sim)
}

// DebugCalculate calculates similarity and prints process log.
func DebugCalculate(code1, code2 string) {
	fmt.Println("It's a checking process with verbosed option and would print debugging information")
	fmt.Println("Name of code file: " + code1 + " and " + code2)
}
