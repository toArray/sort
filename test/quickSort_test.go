package test

import (
	"fmt"
	"luoqiangSort/sort"
	"testing"
)

/**
快排
快排在以下情况下时间复杂度会退化
1.数组已经是正序排过序的
2.数组已经是倒序排过序的
3.所有的元素都相同
*/
func TestQuickSort(t *testing.T) {
	sortData := []int32{9, 2, 11, 4, 65, 8, 1, 1, 9, 9, 3, 10}

	//res := QuickSort(sortData)                                    //递归
	//res = QuickSortSplit(sortData, 0, int32(len(sortData)-1))     //单指针
	res := sort.QuickSortPartition(sortData, 0, int32(len(sortData)-1)) //双指针

	fmt.Println(res)
}
