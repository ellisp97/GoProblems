package Problems

import (
	"sort"
	"strings"
)

/*
Given two strings s1 and s2, we want to visualize how different the two strings are. We will only take into account the lowercase letters (a to z). First let us count the frequency of each lowercase letters in s1 and s2.

s1 = "A aaaa bb c"
s2 = "& aaa bbb c d"

s1 has 4 'a', 2 'b', 1 'c'
s2 has 3 'a', 3 'b', 1 'c', 1 'd'

Example:

s1 = "my&friend&Paul has heavy hats! &"
s2 = "my friend John has many many friends &"
mix(s1, s2) --> "2:nnnnn/1:aaaa/1:hhh/2:mmm/2:yyy/2:dd/2:ff/2:ii/2:rr/=:ee/=:ss"

s1 = "mmmmm m nnnnn y&friend&Paul has heavy hats! &"
s2 = "my frie n d Joh n has ma n y ma n y frie n ds n&"
mix(s1, s2) --> "1:mmmmmm/=:nnnnnn/1:aaaa/1:hhh/2:yyy/2:dd/2:ff/2:ii/2:rr/=:ee/=:ss"


*/

func Mix(s1, s2 string) string {
	alphabase := "abcdefghijklmnopqrstuvwxyz"
	result := []string{}
	for _, c := range alphabase {
		nb_s1 := strings.Count(s1, string(c))
		nb_s2 := strings.Count(s2, string(c))
		if nb_s1 > 1 || nb_s2 > 1 {
			if nb_s1 == nb_s2 {
				result = append(result, "=:"+strings.Repeat(string(c), nb_s1))
			}
			if nb_s1 > nb_s2 {
				result = append(result, "1:"+strings.Repeat(string(c), nb_s1))
			}
			if nb_s1 < nb_s2 {
				result = append(result, "2:"+strings.Repeat(string(c), nb_s2))
			}
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if len(result[i]) == len(result[j]) {
			return result[i] < result[j]
		}
		return len(result[i]) > len(result[j])
	})
	return strings.Join(result, "/")
}
