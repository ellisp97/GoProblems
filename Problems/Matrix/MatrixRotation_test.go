package Problems

import (
	"testing"

	"github.com/ellisp97/GoProblems/Utility"
)

func TestMatrixRotation(t *testing.T) {
	var tests = []struct {
		r             int
		m, reversed_m [][]int
	}{
		{
			2,
			[][]int{{1, 2, 3, 4}, {12, 1, 2, 5}, {11, 4, 3, 6}, {10, 9, 8, 7}},
			[][]int{{3, 4, 5, 6}, {2, 3, 4, 7}, {1, 2, 1, 8}, {12, 11, 10, 9}},
		},
		{
			3,
			[][]int{{1, 1}, {1, 1}},
			[][]int{{1, 1}, {1, 1}},
		},
	}

	for _, c := range tests {
		reversed_m := MatrixRotation(c.m, c.r)
		if !Utility.DoubleArrayEquals(c.reversed_m, reversed_m) {
			t.Errorf("Expected the matrix: %v \n But got the matrix returned: %v", c.reversed_m, reversed_m)
		}
	}
}
