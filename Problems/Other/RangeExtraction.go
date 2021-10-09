package Problems

import (
	"fmt"
	"strconv"
	"strings"
)

/*

A format for expressing an ordered list of integers is to use a comma separated list of
either individual integers or a range of integers denoted by the starting integer separated
from the end integer in the range by a dash, '-'.

The range includes all integers in the interval including both endpoints.

It is not considered a range unless it spans at least 3 numbers. For example "12,13,15-17"
Complete the solution so that it takes a list of integers in increasing order and returns a
correctly formatted string in the range format.

Example:

solution([-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20]);

returns "-6,-3-1,3-5,7-11,14,15,17-20"

*/

// Expected
//     <string>: -6,-3-5,3-8,7-13,14,15,17,17,18
// to equal
//     <string>: -6,-3-1,3-5,7-11,14,15,17-20

func RangeExtraction(arr []int) string {
	seq := []string{}

	if len(arr) == 0 {
		return ""
	}

	temp := []int{}
	temp = append(temp, arr[0])

	for i := 1; i < len(arr); i++ {
		if arr[i] == (arr[i-1] + 1) {
			temp = append(temp, i)
		} else { // If it's not increasing reset temp array, and add as a range if necess
			if len(temp) >= 3 {
				seq = append(seq, fmt.Sprintf("%d-%d", temp[0], temp[len(temp)-1]))
			} else {
				for j := 0; j < len(temp); j++ {
					seq = append(seq, strconv.Itoa(temp[j]))
				}
			}
			temp = []int{arr[i]}
		}
	}

	// add on any remaining
	if len(temp) >= 3 {
		seq = append(seq, fmt.Sprintf("%d-%d", temp[0], temp[len(temp)-1]))
	} else {
		for j := 0; j < len(temp); j++ {
			seq = append(seq, strconv.Itoa(temp[j]))
		}
	}

	return strings.Join(seq, ",")
}
