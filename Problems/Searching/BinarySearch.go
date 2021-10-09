package Searching

import "sort"

/*
	Assumes array is sorted if so O(log n)
	return index position if found
	otherwise return -1
*/

func BinarySearch(arr []int, val, index int) int {

	if len(arr) == 0 || (len(arr) == 1 && arr[0] != val) {
		return -1
	}

	mid := (len(arr) / 2)

	if arr[mid] == val {
		return index + mid
	} else if arr[mid] < val {
		return BinarySearch(arr[mid:], val, index+mid)
	} else {
		return BinarySearch(arr[:mid], val, index)
	}
}

func BinarySearchGoStyle(arr []int, val int) int {
	index := sort.Search(len(arr), func(i int) bool { return arr[i] >= val })

	// e.g. [5,4,3,2,1]
	// reverse_index := sort.Search(len(arr), func(i int) bool { return arr[i] <= val })

	return index
}
