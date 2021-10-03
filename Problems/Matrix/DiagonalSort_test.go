package Problems

import (
	"testing"

	"github.com/ellisp97/GoProblems/Utility"
)

func TestDiagonalMatrixSorted(t *testing.T) {
	var tests = []struct {
		m, m_sorted [][]int
	}{
		{
			[][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}},
			[][]int{{1, 1, 1, 1}, {1, 2, 2, 2}, {1, 2, 3, 3}},
		},
		{
			[][]int{{11, 25, 66, 1, 69, 7}, {23, 55, 17, 45, 15, 52}, {75, 31, 36, 44, 58, 8}, {22, 27, 33, 25, 68, 4}, {84, 28, 14, 11, 5, 50}},
			[][]int{{5, 17, 4, 1, 52, 7}, {11, 11, 25, 45, 8, 69}, {14, 23, 25, 44, 58, 15}, {22, 27, 31, 36, 50, 66}, {84, 28, 75, 33, 55, 68}},
		},
	}

	for _, c := range tests {
		sorted := DiagonalSort(c.m)
		if !Utility.DoubleArrayEquals(sorted, c.m_sorted) {
			t.Errorf("Expected the matrix: %v \n But got the matrix returned: %v", c.m_sorted, sorted)
		}
	}
}
