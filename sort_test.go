package main

import (
	"reflect"
	"testing"
)

func TestSelectionSortAscShouldReturnSortedArray(t *testing.T) {
	//arrange
	arr := []int{5, -2, 23, 34, -5, 26}
	expectedRes := []int{-5, -2, 5, 23, 26, 34}
	//act
	var res = selection_sort_asc(arr)
	//assert
	if reflect.DeepEqual(expectedRes, res) {
		t.Errorf("Expected result is %v but got %v", expectedRes, res)
	}
}

func TestSelectionSortDescShouldReturnSortedArray(t *testing.T) {
	//arrange
	arr := []int{5, -2, 23, 34, -5, 26}
	expectedRes := []int{34, 26, 23, 5, -2, -5}
	//act
	var res = selection_sort_desc(arr)
	//assert
	if reflect.DeepEqual(expectedRes, res) {
		t.Errorf("Expected result is %v but got %v", expectedRes, res)
	}
}
