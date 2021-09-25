package Problems

/*
Consider a sequence u where u is defined as follows:

The number u(0) = 1 is the first one in u.
For each x in u, then y = 2 * x + 1 and z = 3 * x + 1 must be in u too.
There are no other numbers in u.
Example:
u = [1, 3, 4, 7, 9, 10, 13, 15, 19, 21, 22, 27, ...]

1 gives 3 and 4, then 3 gives 7 and 10, 4 gives 9 and 13, then 7 gives 15 and 22 and so on...

Task:
Given parameter n the function dbl_linear (or dblLinear...) returns the element u(n) of the ordered sequence u (ordered with < so there are no duplicates) .

Example:
dbl_linear(10) should return 22

*/

func DblLinear(n int) int {

	arr := []int{1}

	y := func(i int) int {
		return 2*i + 1
	}
	z := func(i int) int {
		return 3*i + 1
	}

	index_y, index_z := 0, 0

	for len(arr) <= n {
		a := y(arr[index_y])
		b := z(arr[index_z])
		if a > b {
			arr = append(arr, b)
			index_z++
		} else if a < b {
			arr = append(arr, a)
			index_y++
		} else {
			arr = append(arr, a)
			index_y++
			index_z++
		}
	}
	return arr[n]
}

// func Contains(a []int, val int) bool{
// 	for _,v := range a{
// 		if v == val{
// 			return true
// 		}
// 	}
// 	return false
// }

// func removeDuplicates(arr []int) []int {
// 	n := len(arr)
// 	j := 1
// 	for i := 1; i < n; i++ {
// 		if arr[i] != arr[i-1] {
// 			arr[j] = arr[i]
// 			j++
// 		}
// 	}
// 	return arr[0:j]
// }
