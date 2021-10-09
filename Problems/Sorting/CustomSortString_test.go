package Sorting

import (
	"testing"
)

func TestCustomSortString(t *testing.T) {
	var tests = []struct {
		order, s string
		res      string
	}{
		{
			"cba",
			"abcd",
			"cbad",
		},
		{
			"cbafg",
			"abcd",
			"cbad",
		},
	}

	for _, c := range tests {
		func_res := CustomSortString(c.order, c.s)
		if c.res != func_res {
			t.Errorf("Error, Expected %v but got %v", c.res, func_res)
		}
	}

}
