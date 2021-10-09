package Problems

import (
	"fmt"
	"sort"
)

// For example, 4 can be partitioned in five distinct ways:

// 4, 3 + 1, 2 + 2, 2 + 1 + 1, 1 + 1 + 1 + 1.

// We can write:

// enum(4) -> [[4],[3,1],[2,2],[2,1,1],[1,1,1,1]] and

// enum(5) -> [[5],[4,1],[3,2],[3,1,1],[2,2,1],[2,1,1,1],[1,1,1,1,1]].

// 1 - n being given (n integer, 1 <= n <= 50) calculate enum(n) ie the partition of n. We will obtain something like that:
// enum(n) -> [[n],[n-1,1],[n-2,2],...,[1,1,...,1]] (order of array and sub-arrays doesn't matter). This part is not tested.

// 2 - For each sub-array of enum(n) calculate its product.
// If n = 5 we'll obtain after removing duplicates and sorting:

// prod(5) -> [1,2,3,4,5,6]
// prod(8) -> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 15, 16, 18]

// 3 - return the range, the average and the median of prod(n) in the following form (example for n = 5):
// "Range: 5 Average: 3.50 Median: 3.50"

func Part(n int) string {

	combinations := make([]int, n+1)
	combinations_content := make([][]int, n+1)

	for i := 0; i < n+1; i++ {
		combinations[i] = 0
	}
	combinations[0] = 1

	for i := 1; i < n+1; i++ {
		for j := 1; j < len(combinations); j++ {
			if j >= i {
				combinations[j] += combinations[j-i]
				combinations_content[j] = append(combinations_content[j], []int{i}...)
				combinations_content[j] = append(combinations_content[j], combinations_content[j-i]...)
			}
		}
	}

	result := combinations_content[1:]
	// for each of these values we need to sort and remove duplicate
	res := removeDuplicatesAndSort(result)

	range_value := calculateRange(res)
	average_value := calculateAverage(res)
	median_value := calculateMedian(res)

	fmt.Println(res)
	return fmt.Sprintf("Range: %d Average: %.2f Median: %.2f", range_value, average_value, median_value)
}

func calculateSum(arr [][]int) []int {
	res := []int{}
	for i := 0; i < len(arr); i++ {
		if len(arr) == 0 {
			continue
		}
		multiple_value := arr[i][0]
		for j := 1; j < len(arr[i]); j++ {
			multiple_value *= arr[i][j]
		}
		res = append(res, multiple_value)
	}
	return res
}

func calculateMedian(arr []int) float64 {
	v := 0.0
	fmt.Println(arr)
	mid := len(arr) / 2
	if len(arr)%2 == 0 {
		v = float64(float64(arr[mid-1])+float64(arr[mid])) / float64(2)
	} else {
		v = float64(arr[mid])
	}
	return v

}

func calculateRange(arr []int) int {
	v := 0

	if len(arr) > 1 {
		v = arr[len(arr)-1] - arr[0]
	}

	return v
}

func calculateAverage(arr []int) float64 {
	v := 0.0
	for i := 0; i < len(arr); i++ {
		v += float64(arr[i])
	}
	if len(arr) > 0 {
		v /= float64(len(arr))
	}

	return v
}

func removeDuplicatesAndSort(arr [][]int) []int {
	li := []int{}
	res := map[int]bool{}
	for _, v := range arr {
		for _, v1 := range v {
			if _, ok := res[v1]; !ok {
				res[v1] = true
				li = append(li, v1)
			}
		}
	}
	sort.Ints(li)
	return li
}
