package Problems

func reverse(x int) int {
	INT_MIN, INT_MAX := -2147483648, 2147483647
	res := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if res > INT_MAX/10 {
			return 0
		}
		if res < INT_MIN/10 {
			return 0
		}
		res = res*10 + pop
	}
	return res
}
