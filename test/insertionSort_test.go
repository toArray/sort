package test

import (
	"fmt"
	"luoqiangSort/sort"
	"testing"
)

//ζε₯ζεΊ
func TestInsertionSort(t *testing.T) {
	sortData := []int32{9, 2, 11, 4, 65, 8, 1, 1, 9, 9, 3}

	res := sort.InsertionSort(sortData)

	fmt.Println(res)
}
