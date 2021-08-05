package sort

//时间复杂度:平均:O(n^2)、最好:O(n)、最坏:O(n^2)
//空间复杂度:O(1)
//稳定性:稳定

/*
BubbleSort
@Desc:冒泡排序
@Param:arr 	[]int32		待排数据
@Return:	[]int32		排序结果
*/
func BubbleSort(arr []int32) []int32 {
	l := len(arr)
	if l < 2 {
		return arr
	}

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}
