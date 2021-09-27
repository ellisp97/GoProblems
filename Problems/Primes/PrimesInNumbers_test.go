package Problems

import (
	"testing"
)

func TestPrimeInNumbers(t *testing.T) {
	var tests = []struct {
		n   int
		res string
	}{
		{
			7775460,
			"(2**2)(3**3)(5)(7)(11**2)(17)",
		},
		{
			79340,
			"(2**2)(5)(3967)",
		},
	}

	for _, c := range tests {
		func_result := GetPrimeFactors(c.n)
		if c.res != func_result {
			t.Errorf("Expected %s, but got %s", c.res, func_result)
		}
	}
}
