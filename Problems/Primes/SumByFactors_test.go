package Problems

import (
	"testing"
)

func TestSumByFactors(t *testing.T) {
	var tests = []struct {
		arr []int
		res string
	}{
		{
			[]int{12, 15},
			"(2 12)(3 27)(5 15)",
		},
		{
			[]int{15, 21, 24, 30, 45},
			"(2 54)(3 135)(5 90)(7 21)",
		},
	}

	for _, c := range tests {
		func_result := SumByFactors(c.arr)
		if c.res != func_result {
			t.Errorf("Expected %s, but got %s", c.res, func_result)
		}
	}
}
