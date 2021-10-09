package Searching

import "testing"

func TestBinarySearches(t *testing.T) {
	var tests = []struct {
		arr []int
		val int
		pos int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			5,
			4,
		},
		{
			[]int{1, 2, 3},
			3,
			2,
		},
		{
			[]int{0},
			1,
			-1,
		},
	}

	for _, c := range tests {
		bin_pos := BinarySearch(c.arr, c.val, 0)
		if bin_pos != c.pos {
			t.Errorf("Expected position %d but got %d", c.pos, bin_pos)
		}
	}

}
