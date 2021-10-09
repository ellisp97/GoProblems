package Problems

func CheckArrays(arr1, arr2 []int) bool {

	if arr1 == nil || arr2 == nil || len(arr1) != len(arr2) {
		return false
	}

	if len(arr1) == 0 {
		return true
	}

	hmap := make(map[int]int)

	for _, val := range arr1 {
		hmap[val*val]++
	}

	for _, val := range arr2 {
		if v, ok := hmap[val]; !ok || v == 0 {
			return false
		}
		hmap[val]--
	}

	return true
}
