package Problems

// Given two strings s and goal, return true if you can swap two letters in s so the result is equal to goal, otherwise, return false.

// Swapping letters is defined as taking two indices i and j (0-indexed) such that i != j and swapping the characters at s[i] and s[j].

// For example, swapping at indices 0 and 2 in "abcd" results in "cbad".

// Example 1:

// Input: s = "ab", goal = "ba"
// Output: true
// Explanation: You can swap s[0] = 'a' and s[1] = 'b' to get "ba", which is equal to goal.
// Example 2:

// Input: s = "ab", goal = "ab"
// Output: false
// Explanation: The only letters you can swap are s[0] = 'a' and s[1] = 'b', which results in "ba" != goal.

type Tuple struct {
	val1 rune
	val2 rune
}

func (t Tuple) reverse() Tuple {
	t.val1, t.val2 = t.val2, t.val1
	return t
}

func buddyString(A, B string) bool {
	if len(A) != len(B) {
		return false
	}

	// We need to check that there is a valid swap which can be done
	if A == B {
		hmap := make(map[rune]struct{})
		for _, v := range A {
			if _, ok := hmap[v]; ok {
				return true
			}
			hmap[v] = struct{}{}
		}
	}

	// If theyre not equal keep track of pairs of differences
	goal_arr := []rune(B)
	diff := []Tuple{}
	for i, v := range A {
		if v != goal_arr[i] {
			diff = append(diff, Tuple{val1: v, val2: goal_arr[i]})
		}
		if len(diff) >= 3 {
			return false
		}
	}

	return len(diff) == 2 && diff[0] == diff[1].reverse()
}
