package main

import (
	"math"
)

func selection_sort_desc(arr []int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		maxIndex := math.MinInt
		for j := i; j < len(arr); j++ {
			if arr[j] > maxIndex {
				maxIndex = j
			}
		}
		arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
	}
	return arr
}

func selection_sort_asc(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIndex := math.MaxInt
		for j := i; j < len(arr); j++ {
			if arr[j] < minIndex {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}
