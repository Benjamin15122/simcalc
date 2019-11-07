package simcalc

import (
	"testing"
)

func TestSimCalculateOnLcs(t *testing.T) {
	e := SimCalculate("testcase/lcs/lcs1.cpp", "testcase/lcs/lcs2.cpp")
	if e != nil {
		t.Error("SimCalculate lcs: Failed")
	} else {
		t.Log("SimCalculate lcs: Passed")
	}
}
