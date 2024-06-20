//go:build ignored

package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	if Sum(5, 5) != 10 {
		t.Error("Failed to calculate correct sum")
	}
}
