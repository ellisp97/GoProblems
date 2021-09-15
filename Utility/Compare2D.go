package Utility

func DoubleArrayEquals(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !IntArrayEquals(v, b[i]) {
			return false
		}
	}
	return true
}

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
