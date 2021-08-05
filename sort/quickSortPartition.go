package sort

/*
时间复杂度:平均:O(nlogn)、最好:O(nlogn)、最坏:O(n^2)
空间复杂度:O(nlogn)
稳定性:不稳定
原理:
	1.设定一个基准元素
	2.设置最左边,最右边2个指针
	3.右边大于基准元素,指针左移。左边小于基准元素,指针右移。2个都移动完后交换
	4.交换基准元素
*/

/*
QuickSortPartition
快排 - 双指针
@Param:arr 	[]int32		待排数据
@Param:low 	int32		低位索引
@Param:high int32		高位索引
@Return:	[]int32		排序结果
*/
func QuickSortPartition(arr []int32, low int32, high int32) []int32 {
	l := len(arr)
	if l < 2 {
		return arr
	}

	if low < high {
		index := partition(arr, low, high)
		QuickSortPartition(arr, low, index-1)
		QuickSortPartition(arr, index+1, high)
	}

	return arr
}

/*
partition
@Desc:快排双指针处理
@Param:arr 			[]int32		待排数据
@Param:startIndex 	int32		低位索引
@Param:endIndex 	int32		高位索引
@Return:			int32		基准元素索引
*/
func partition(arr []int32, startIndex, endIndex int32) int32 {
	pivot := arr[startIndex] //基准元素
	left := startIndex       //第一个指针左边
	right := endIndex        //第二个指针右边

	//左右不相等就得继续排
	for left != right {
		//right大于基准元素,right左移一位,如果right元素小于pivot元素的话停止移动
		for left < right && pivot < arr[right] {
			right--
		}

		//left小于基准元素,left右移一位,如果left元素大于pivot元素,停止移动
		for left < right && pivot >= arr[left] {
			left++
		}

		//左右都停止移动后,交换
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	//交换基准元素
	arr[startIndex], arr[left] = arr[left], arr[startIndex]

	return left
}
