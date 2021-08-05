package sort

/*
时间复杂度:平均:O(n^1.3) 最好:O(n) 最坏:O(n^2)
空间复杂度:O(1)
稳定性:不稳定
原理:
	1.设置在增量flag进行多个分组
	2.分别对多组进行插入排序
*/

/*
ShellSort
希尔排序
@Param:arr 	[]int32		待排数据
@Return:	[]int32		排序结果
*/
func ShellSort(arr []int32) []int32 {
	l := len(arr)
	if l < 2 {
		return arr
	}

	flag := l / 2

	for flag > 0 {
		for i := flag; i < l; i++ {
			for j := i; j-flag >= 0 && arr[j] < arr[j-flag]; j -= flag {
				arr[j], arr[j-flag] = arr[j-flag], arr[j]
			}
		}
		flag /= 2
	}

	return arr
}
