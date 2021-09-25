package Problems

import (
	"testing"

	"github.com/ellisp97/GoProblems/Utility"
)

func TestSquaresIntoSquares(t *testing.T) {

	var tests = []struct {
		n   int
		arr []int
	}{
		{
			5,
			[]int{3, 4},
		},
		{
			50,
			[]int{1, 3, 5, 8, 49},
		},
		{
			2,
			[]int{},
		},
		{
			7,
			[]int{2, 3, 6},
		},
	}

	for _, c := range tests {
		res := Decompose(c.n)
		if !Utility.IntArrayEquals(res, c.arr) {
			t.Errorf("Expected %v but got %v", c.arr, res)
		}
	}

}
