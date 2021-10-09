package Problems

func isValid(s string) bool {
	stack := []rune{}
	hmap := make(map[rune]rune)
	hmap['}'] = '{'
	hmap[')'] = '('
	hmap[']'] = '['
	for _, v := range s {
		if len(stack) > 0 && stack[len(stack)-1] == hmap[v] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}
