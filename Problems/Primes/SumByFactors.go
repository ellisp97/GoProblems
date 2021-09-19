package Problems

import (
	"fmt"
	"sort"
)

/*
Given an array of positive or negative integers

I= [i1,..,in]

you have to produce a sorted array P of the form

[ [p, sum of all ij of I for which p is a prime factor (p positive) of ij] ...]

P will be sorted by increasing order of the prime numbers.

Example:
I = (/12, 15/); // result = "(2 12)(3 27)(5 15)"
*/

func SumByFactors(arr []int) string {
	var sums = make(map[int]int)
	res := ""

	for _, val := range arr {
		pos_val := val
		if val < 0 {
			pos_val = val * -1
		}

		for i := 2; i*i <= pos_val; i++ {
			if pos_val%i == 0 {
				sums[i] += val

				for pos_val%i == 0 {
					pos_val /= i
				}
			}
		}

		if pos_val > 1 {
			sums[pos_val] += val
		}
	}

	keys := []int{}
	for key := range sums {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		res += fmt.Sprintf("(%d %d)", key, sums[key])
	}

	fmt.Println(sums)
	return res
}
