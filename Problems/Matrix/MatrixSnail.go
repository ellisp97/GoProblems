package Problems

import (
	"fmt"
)

/*
Q: You are given a 2D matrix of dimension travel through the matrix
in a snail fashion so elements are returned in a list

  --->
  |> |
  <---

Ex:

array = [[1,2,3],
         [8,9,4],
         [7,6,5]]

snail(array) #=> [1,2,3,4,5,6,7,8,9]

[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
*/

func SnailMatrix(matrix [][]int) []int {
	count := len(matrix[0]) * len(matrix)
	arr := []int{}
	index := 0

	x_min, x_max := 0, len(matrix[0])-1
	y_min, y_max := 0, len(matrix)-1
	for index < count {

		for i := x_min; i <= x_max; i++ {
			arr = append(arr, matrix[y_min][i])
			index++
		}
		y_min++

		for j := y_min; j <= y_max; j++ {
			arr = append(arr, matrix[j][x_max])
			index++
		}
		x_max--

		for i := x_max; i >= x_min; i-- {
			arr = append(arr, matrix[y_max][i])
			index++
		}
		y_max--
		fmt.Println(arr)
		for i := y_max; i >= y_min; i-- {
			fmt.Printf("matrix %v x: %v, y: %v\n", matrix, x_min, i)
			arr = append(arr, matrix[i][x_min])
			index++
		}
		x_min++

	}
	return arr
}
