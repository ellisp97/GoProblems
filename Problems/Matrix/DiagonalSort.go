package Problems

import "sort"

type Queue struct {
	arr []int
}

func (q *Queue) dequeue() (ret int) {
	ret = q.arr[0]
	q.arr = q.arr[1:]
	return
}

func (q *Queue) enqueue(val int) {
	q.arr = append(q.arr, val)
}

func DiagonalSort(mat [][]int) [][]int {

	// 1. Access each diagonal
	// 2. Store each diagonal
	// 3. Sort each diagonal
	n, m := len(mat[0]), len(mat)

	// diagonals are indexed by the difference of their row and columns
	diagonals := make(map[int]*Queue)

	// 1.
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if _, ok := diagonals[row-col]; !ok { // 2.
				diagonals[row-col] = &Queue{arr: []int{mat[row][col]}}
			} else {
				diagonals[row-col].enqueue(mat[row][col])
			}
		}
	}

	// store each value in our diagonal hash map
	for k, _ := range diagonals {
		sort.Ints(diagonals[k].arr)
	}

	// Loop back through matrix adding diagonals back in
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			mat[row][col] = diagonals[row-col].dequeue()
		}
	}

	return mat
}
