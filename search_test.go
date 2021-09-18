package main

import "testing"

func TestBinarySearchShouldReturnIndex(t *testing.T) {
	//arrange
	arr := []int{1, 2, 3, 4, 5, 6}
	search := 5
	expectedRes := 4
	//act
	var res = binary_search(arr, search)
	//assert
	if res != expectedRes {
		t.Errorf("Expected result is %v but got %v", expectedRes, res)
	}
}

func TestBinarySearchShouldReturnNotFound(t *testing.T) {
	//arrange
	arr := []int{1, 2, 3, 5, 6}
	search := 4
	expectedRes := -1
	//act
	var res = binary_search(arr, search)
	//assert
	if res != expectedRes {
		t.Errorf("Expected result is %v but got %v", expectedRes, res)
	}
}
