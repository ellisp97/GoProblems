package Sorting

import (
	"testing"

	"github.com/ellisp97/GoProblems/Utility"
)

func TestSort(t *testing.T) {
	var tests = []struct {
		arr        []int
		sorted_arr []int
	}{
		{
			[]int{1, 7, 3, 2, 5},
			[]int{1, 2, 3, 5, 7},
		},
		{
			[]int{1, 2, 3, 5, 6, 4},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1},
			[]int{1},
		},
	}

	for _, c := range tests {
		heap_arr := HeapSort(c.arr)
		bubble_arr := BubbleSort(c.arr)

		if !Utility.IntArrayEquals(heap_arr, c.sorted_arr) {
			t.Errorf("Error in Heap Sort\n Expected %v but got %v", c.sorted_arr, heap_arr)
		}
		if !Utility.IntArrayEquals(bubble_arr, c.sorted_arr) {
			t.Errorf("Error in Bubble Sort\n Expected %v but got %v", c.sorted_arr, bubble_arr)
		}

	}

}
