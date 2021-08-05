package sort

import (
	"sync"
)

//时间复杂度:平均:O(nlogn)、最好:O(nlogn)、最坏:O(n^2)
//空间复杂度:O(nlogn)
//稳定性:不稳定

/*
QuickSort
快排 - 递归
@Param:arr 	[]int32		待排数据
@Return: 	[]int32		排序结果
*/
func QuickSort(arr []int32) []int32 {
	l := len(arr)
	if l < 2 {
		return arr
	}

	index := arr[0]
	left := make([]int32, 0)
	right := make([]int32, 0)

	j := l
	for i := 1; i < j; i++ {
		if arr[i] <= index {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}

		if j != l {
			if arr[i] <= index {
				left = append(left, arr[j])
			} else {
				right = append(right, arr[j])
			}
			j--
		}
	}

	wait := sync.WaitGroup{}
	wait.Add(2)
	go func() {
		left = QuickSort(left)
		wait.Done()
	}()

	go func() {
		right = QuickSort(right)
		wait.Done()
	}()

	wait.Wait()
	res := make([]int32, 0)
	res = append(res, left...)
	res = append(res, index)
	res = append(res, right...)

	return res
}
