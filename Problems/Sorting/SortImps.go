package Sorting

func BubbleSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
	return data
}

func HeapSort(arr []int) []int {
	length := len(arr)
	x := length/2 - 1 // Last parent node in heap
	y := length - 1   // Last child of heap

	// 1. Create max heap starting at last parent node working up
	for x >= 0 {
		heapify(arr, length, x)
		x--
	}

	// 2). Switch the root (or the top of the tree) with the last node and then remove it from the heap.
	for y >= 0 {
		arr[0], arr[y] = arr[y], arr[0]
		// 3). Rebuild the max heap again and repeat until there is only one element remaining.
		heapify(arr, y, 0)
		y--
	}
	return arr
}

func heapify(arr []int, length, i int) []int {
	largest := i
	left := i*2 + 1
	right := left + 1

	if left < length && arr[left] > arr[largest] {
		largest = left
	}

	if right < length && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, length, largest)
	}
	return arr
}
