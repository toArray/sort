package sort

//时间复杂度:平均:O(nlog2n)、最好:O(nlog2n)、最坏:O(nlog2n)
//空间复杂度:O(n)
//稳定性:稳定

/*
MergeSort
@Desc:归并排序
@Param:arr 	[]int32		待排数据
@Return: 	[]int32		排序结果
*/
func MergeSort(arr []int32) []int32 {
	l := len(arr)
	if l < 2 {
		return arr
	}

	middle := l / 2
	left := arr[:middle]
	right := arr[middle:]

	left = MergeSort(left)
	right = MergeSort(right)
	return merge(left, right)
}

//合并
func merge(left, right []int32) []int32 {
	l, r := 0, 0
	var temp []int32

	for {
		//如果左边已经结束,将右边剩余直接添加
		if l >= len(left) {
			temp = append(temp, right[r:]...)
			break
		}

		//如果右边已经结束,将左边剩余直接添加
		if r >= len(right) {
			temp = append(temp, left[l:]...)
			break
		}

		//左右相比,将较小元素添加
		if left[l] < right[r] {
			temp = append(temp, left[l])
			l++
		} else {
			temp = append(temp, right[r])
			r++
		}
	}

	return temp
}
