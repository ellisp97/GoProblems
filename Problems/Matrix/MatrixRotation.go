package Problems

import (
	"fmt"
	"math"
)

/*
Q: You are given a 2D matrix of dimension  and a positive integer . You have to rotate the matrix  times and print the resultant matrix. Rotation should be in anti-clockwise direction.


Ex:

It is guaranteed that the minimum of m and n will be even.

As an example rotate the Start matrix by 2:

Start         First           Second
 1 2 3 4        2  3  4  5      3  4  5  6
12 1 2 5  ->   1  2  3  6 ->   2  3  4  7
11 4 3 6      12  1  4  7       1  2  1  8
10 9 8 7      11 10  9  8     12 11 10  9

*/

func MatrixRotation(matrix [][]int, r int) [][]int {
	// Obtain outer layer and shift the array as one list
	temp_list := [][]int{}

	n, m := len(matrix[0]), len(matrix)
	layers := int(math.Min(float64(n), float64(m))) / 2 // Compute no of layers

	for i := 0; i < layers; i++ {
		temp := []int{}
		// Top Layer
		for j := i; j < n-1-i; j++ {
			temp = append(temp, matrix[i][j])
		}
		// Side Layer
		for j := i; j < m-1-i; j++ {
			temp = append(temp, matrix[j][n-1-i])
		}
		//Reverse Order
		for j := n - 1 - i; j > i; j-- {
			temp = append(temp, matrix[m-1-i][j])
		}
		//Upwards
		for j := m - 1 - i; j > i; j-- {
			temp = append(temp, matrix[j][i])
		}
		temp_list = append(temp_list, temp)
	}

	for i := 0; i < layers; i++ {
		row := temp_list[i]
		index := r % len(row)
		// Top Layer
		for j := i; j < n-1-i; j++ {
			matrix[i][j] = row[index]
			index++
			index %= len(row)
		}
		// Side Layer
		for j := i; j < m-1-i; j++ {
			matrix[j][n-1-i] = row[index]
			index++
			index %= len(row)
		}
		//Reverse Order
		for j := n - 1 - i; j > i; j-- {
			matrix[m-1-i][j] = row[index]
			index++
			index %= len(row)
		}
		//Upwards
		for j := m - 1 - i; j > i; j-- {
			matrix[j][i] = row[index]
			index++
			index %= len(row)
		}
	}
	fmt.Printf("Rotated Matrix: %+v \n", matrix)
	return matrix
}
