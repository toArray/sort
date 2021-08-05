package sort

/*
时间复杂度:平均:O(n^2)、最好:O(n^2)、最坏:O(n^2)
空间复杂度:O(1)
稳定性:不稳定
原理:
	1.在一个长度为 N 的无序数组中，第一次遍历 n-1 个数找到最小的和第一个数交换。
	2.第二次从下一个数开始遍历 n-2 个数，找到最小的数和第二个数交换。
	3.重复以上操作直到第 n-1 次遍历最小的数和第 n-1 个数交换，排序完成。
*/

/*
SelectionSort
@Desc:选择排序
@Param:arr 	[]int32		待排数据
@Return:	[]int32		排序结果
*/
func SelectionSort(arr []int32) []int32 {
	l := len(arr)
	if l < 2 {
		return arr
	}

	for i := 0; i < l-1; i++ {
		//从以下找出最小索引
		minIndex := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[i] && arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr
}
