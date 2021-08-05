package sort

//时间复杂度:平均:O(n^2)、最好:O(n)、最坏:O(n^2)
//空间复杂度:O(1)
//稳定性:稳定

/*
InsertionSort
@Desc:插入排序
@Param:arr 	[]int32		待排数据
@Return:	[]int32		排序结果
*/
func InsertionSort(arr []int32) []int32 {
	l := len(arr)
	if l < 2 {
		return arr
	}

	for i := 1; i < l; i++ {
		preIndex := i - 1
		curValue := arr[i]
		for preIndex >= 0 && arr[preIndex] > curValue {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = curValue
	}

	return arr
}

/*
InsertSortFloat64
@Desc:插入排序float64
@Param:data []float64	待排序数据
@Return:res []float64	排序结果
*/
func InsertSortFloat64(arr []float64) (res []float64) {
	if len(arr) <= 1 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		preIndex := i - 1
		curValue := arr[i]
		for preIndex >= 0 && curValue <= arr[preIndex] {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = curValue
	}

	return arr
}
