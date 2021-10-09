package Sorting

/*
You are given two strings order and s. All the words of order are unique and were sorted in some custom order previously.
Permute the characters of s so that they match the order that order was sorted. More specifically, if a character x occurs before a character y in order, then x should occur before y in the permuted string.
Return any permutation of s that satisfies this property.

Example 1:

Input: order = "cba", s = "abcd"
Output: "cbad"
Explanation:
"a", "b", "c" appear in order, so the order of "a", "b", "c" should be "c", "b", and "a".
Since "d" does not appear in order, it can be at any position in the returned string. "dcba", "cdba", "cbda" are also valid outputs.
*/

import "strings"

func CustomSortString(order string, s string) string {

	var sb strings.Builder
	hmap := map[rune]int{}
	for _, v := range s {
		hmap[v]++
	}

	// write all chars according to the order in order, and the count in s
	for _, v := range order {
		sb.WriteString(strings.Repeat(string(v), hmap[v]))
		hmap[v] = 0
	}

	// write any remaining characters
	for k, v := range hmap {
		sb.WriteString(strings.Repeat(string(k), v))
	}
	return sb.String()
}
