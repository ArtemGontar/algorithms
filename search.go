package main

func binary_search(arr []int, search int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		middle := (left + right) / 2
		if search > arr[middle] {
			left = middle + 1
		}
		if search < arr[middle] {
			right = middle - 1
		}
		if search == arr[middle] {
			return middle
		}
	}
	return -1
}
