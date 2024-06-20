//go:build ignored
package main

import "testing"

//Testing function Mult(...)
func TestMult(t *testing.T) {
	if res := Mult(1, 1, 1, 1); res != 1 {
		t.Error("Exepcted ", 1, "Actual ", res)
	}

	if res := Mult(0, 1, 1, 1); res != 0 {
		t.Error("Exepcted ", 0, "Actual ", res)
	}

	if res := Mult(10, 10, 10, 10); res != 10000 {
		t.Error("Exepcted ", 10000, "Actual ", res)
	}
}
