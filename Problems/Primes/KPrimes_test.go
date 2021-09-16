package Problems

import (
	"testing"

	"github.com/ellisp97/GoProblems/Utility"
)

func TestConsecKPrimes(t *testing.T) {
	var tests = []struct {
		k, start, end int
		res           []int
	}{
		{
			5,
			500,
			600,
			[]int{500, 520, 552, 567, 588, 592, 594},
		},
	}

	for _, c := range tests {
		func_result := CountKPrimes(c.k, c.start, c.end)
		if !Utility.IntArrayEquals(c.res, func_result) {
			t.Errorf("Expected %v, but got %v", c.res, func_result)
		}
	}
}
