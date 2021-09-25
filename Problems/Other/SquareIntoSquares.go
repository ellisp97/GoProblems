package Problems

/*
Given a positive integral number n,
return a strictly increasing sequence array of
numbers, so that the sum of the squares is
equal to n².

Examples
decompose(11) must return [1,2,4,10].
Note that there are actually two ways
to decompose 11², 11² = 121 = 1 + 4 + 16 + 100 = 1² + 2² + 4² + 10²
but don't return [2,6,9], since 9 is smaller than 10.
*/

func Decompose(n int) []int {
	arr := loop(n*n, n)
	return arr
}

func loop(sq, i int) []int {
	if sq < 0 {
		return nil
	}
	if sq == 0 {
		return []int{}
	}

	for j := i - 1; j > 0; j-- {
		remainder := loop(sq-j*j, j)
		if remainder != nil {
			return append(remainder, j)
		}
	}
	return nil
}
