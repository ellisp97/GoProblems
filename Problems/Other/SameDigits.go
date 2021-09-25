package Problems

import (
	"math"
	"sort"

	"github.com/ellisp97/GoProblems/Utility"
)

/*

Create a function that takes a positive integer and returns the next bigger number that can be formed by rearranging its digits. For example:

12 ==> 21
513 ==> 531
2017 ==> 2071
nextBigger(num: 12)   // returns 21
nextBigger(num: 513)  // returns 531
nextBigger(num: 2017) // returns 2071
If the digits can't be rearranged to form a bigger number, return -1 (or nil in Swift):

9 ==> -1
111 ==> -1
531 ==> -1
nextBigger(num: 9)   // returns nil
nextBigger(num: 111) // returns nil
nextBigger(num: 531) // returns nil

*/

func NextBigger(n int) int {
	arr := IntToSlice([]int{}, n)

	// Maxium value to check
	sort.Slice(arr, func(i, j int) bool {
		return arr[j] < arr[i]
	})

	num := n + 1
	max := SliceToInt(arr)

	for num <= max {
		if NumsHaveSameDigits(num, n) {
			return num
		}
		num++
	}

	return -1
}

func NumsHaveSameDigits(x, y int) bool {
	arr_x := IntToSlice([]int{}, x)
	arr_y := IntToSlice([]int{}, y)
	sort.Ints(arr_x)
	sort.Ints(arr_y)

	return Utility.IntArrayEquals(arr_x, arr_y)
}

func SliceToInt(arr []int) int {
	sum := 0
	power := 0
	for i := len(arr) - 1; i >= 0; i-- {
		sum += arr[i] * int(math.Pow(10, float64(power)))
		power++
	}
	return sum
}

func IntToSlice(arr []int, n int) []int {
	if n != 0 {
		i := n % 10
		arr = append([]int{i}, arr...)
		return IntToSlice(arr, n/10)
	}
	return arr
}
