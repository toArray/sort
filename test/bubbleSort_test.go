package test

import (
	"fmt"
	"luoqiangSort/sort"
	"testing"
)

//冒泡排序
func TestBubbleSort(t *testing.T) {
	sortData := []int32{9, 2, 11, 4, 65, 8, 1, 1, 9, 9, 3, 43, 5, 65, 1, 68, 0}

	res := sort.BubbleSort(sortData)

	fmt.Println(res)
}
