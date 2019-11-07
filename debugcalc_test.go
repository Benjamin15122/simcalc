package simcalc

import (
	"testing"
)

func TestDebugCalculateOnLcs(t *testing.T) {
	e := DebugCalculate("testcase/lcs/lcs1.cpp", "testcase/lcs/lcs2.cpp")
	if e != nil {
		t.Error("Testcase lcs: Failed")
	} else {
		t.Log("Testcase lcs: Passed")
	}
}
