package test

import (
	"fmt"
	"luoqiangSort/sort"
	"testing"
)

//选择排序
func TestSelectionSort(t *testing.T) {
	sortData := []int32{9, 2, 11, 4, 65, 8, 1, 1, 9, 9, 3, 3, 4, 5, 76, 7, 1, 1}

	res := sort.SelectionSort(sortData)

	fmt.Println(res)
}
