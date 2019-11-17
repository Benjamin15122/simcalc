package simcalc

import (
	"fmt"
	"testing"
)

func TestLCSComputer(t *testing.T) {
	len, value := LCSComputer("AGGTABTABTABTAB", "GXTXAYBTABTABTAB")
	if len < 13 {
		t.Errorf("Fail with %d %x", len, value)
	}
}

func TestFSGenerator(t *testing.T) {
	s, err := FSGenerator("test.ll")
	fmt.Print(s)
	fmt.Println(err)
}

func TestIRGenerator(t *testing.T) {
	IRGenerator("testcase/lcs/lcs1.cpp", "test.ll")
}

func TestSimCalculateOnLcs(t *testing.T) {
	SimCalculate("testcase/lcs/lcs1.cpp", "testcase/lcs/lcs2.cpp")
}

func TestDebugCalculateOnLcs(t *testing.T) {
	DebugCalculate("testcase/lcs/lcs1.cpp", "testcase/lcs/lcs2.cpp")
}
