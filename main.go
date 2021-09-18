package main

import (
	"fmt"
)

func main() {
	//fmt.Println(binary_search([]int{1, 3, 4, 5, 6}, 4))

	fmt.Println(selection_sort_asc([]int{1, 3, 4, 5, 6}))
	fmt.Println(selection_sort_desc([]int{1, 3, 4, 5, 6}))
}
