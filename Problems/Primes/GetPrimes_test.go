package Problems

import (
	"testing"

	"github.com/ellisp97/GoProblems/Utility"
)

func TestSieve(t *testing.T) {
	var tests = []struct {
		n      int
		result []bool
	}{
		{
			5,
			[]bool{false, false, true, true, false, true},
		},
		{
			10,
			[]bool{false, false, true, true, false, true, false, true, false, false, false},
		},
		{
			0,
			[]bool{},
		},
	}

	for _, c := range tests {
		res := SieveOfEratosthenes(c.n)
		if !Utility.BoolArrayEquals(res, c.result) {
			t.Errorf("Expected the bool arr %v but instead got %v \n", c.result, res)
		}
	}
}
