package sort

/*
时间复杂度:平均:O(nlogn)、最好:O(nlogn)、最坏:O(n^2)
空间复杂度:O(nlogn)
稳定性:不稳定
原理:
	1.设定一个基准元素
	2.将数组第一个数组元素设置为比较元素
	3.从数组第二个开始遍历比较,若找到比基准大的则跳过,否则将其与前面较大的进行交换
	4.将刚开始的基准数放到中间
	5.循环234
*/

/*
QuickSortSplit
快排 - 单指针
@Param:arr 	[]int32		待排数据
@Param:low 	int32		低位索引
@Param:high int32		高位索引
@Return:	[]int32		排序结果
*/
func QuickSortSplit(arr []int32, low int32, high int32) []int32 {
	l := len(arr)
	if l < 2 {
		return arr
	}

	if low < high {
		index := split(arr, low, high)
		QuickSortSplit(arr, low, index-1)
		QuickSortSplit(arr, index+1, high)
	}

	return arr
}

/*
split
@Desc:单指针处理
@Param:arr 	[]int32		待排数据
@Param:low 	int32		低位索引
@Param:high int32		高位索引
@Return:	int32		基准元素索引
*/
func split(arr []int32, low int32, high int32) int32 {
	index := low          //基准元素下标
	compare := arr[index] //将数组第一个数组元素设置为比较元素

	//单边扫描,基准元素后一位开始比较
	for i := low + 1; i <= high; i++ {
		if arr[i] < compare {
			index++
			arr[index], arr[i] = arr[i], arr[index] //交换2个比较元素
		}
	}

	//把基准元素和第一个交换,放到中间来
	arr[low], arr[index] = arr[index], arr[low]

	return index
}
