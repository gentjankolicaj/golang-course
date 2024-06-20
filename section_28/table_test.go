//go:build ignored

package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	type test struct {
		data   []int
		result int
	}

	var testData []test = []test{
		test{data: []int{1, 2, 3, 4, 4}, result: 14},
		test{data: []int{2, 2, 2, 2, 2}, result: 10},
		test{data: []int{-3, -3, -3, -3}, result: -12},
	}

	for i := 0; i < len(testData); i++ {
		testStuct := testData[i]
		sum := Sum(testStuct.data...)
		if sum != testStuct.result {
			t.Error("Expected ", testStuct.result, " Actual ", sum)
		}

	}

}
