package Problems

import (
	"testing"

	"github.com/ellisp97/GoProblems/Utility"
)

func TestSnailMatrix(t *testing.T) {
	var tests = []struct {
		matrix [][]int
		arr    []int
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
	}
	for _, c := range tests {
		res := SnailMatrix(c.matrix)
		if !Utility.IntArrayEquals(res, c.arr) {
			t.Errorf("Expected the snail array to be %v but instead got %v", c.arr, res)
		}
	}
}
