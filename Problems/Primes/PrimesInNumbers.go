package Problems

import (
	"fmt"
	"sort"
	"strings"
)

// Given a positive number n > 1 find the prime factor decomposition of n. The result will be a string with the following form :

//  "(p1**n1)(p2**n2)...(pk**nk)"
// with the p(i) in increasing order and n(i) empty if n(i) is 1.

// Example: n = 86240 should return "(2**5)(5)(7**2)(11)"

func GetPrimeFactors(n int) string {

	hmap := make(map[int]int)

	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			n /= i
			hmap[i]++
		}
	}

	if n > 1 {
		hmap[n]++
	}

	keys := make([]int, len(hmap))

	for key, _ := range hmap {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	var sb strings.Builder
	for _, val := range keys {
		if hmap[val] > 1 {
			sb.WriteString(fmt.Sprintf("(%d**%d)", val, hmap[val]))
		} else if hmap[val] == 1 {
			sb.WriteString(fmt.Sprintf("(%d)", val))
		}
	}

	return sb.String()
}
