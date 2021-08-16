package sort

//时间复杂度:平均:O(n+k)、最好:O(n+k)、最坏:O(n+k)
//空间复杂度:O(n+k)
//稳定性:不稳定

/*
CountingSortOptimize
@Desc:计数排序 - 优化版本
@Param:arr 	[]int32		待排数据
@Return:res []int32		排序结果
*/
func CountingSortOptimize(arr []int32) (res []int32) {
	if len(arr) == 0 {
		return
	}

	//获得最大值并开始记录
	maxNum, minNum := getMaxMinNum(arr)
	lenNum := maxNum - minNum + 1
	countSlice := make([]int32, lenNum)
	for _, value := range arr {
		countSlice[value-minNum] += 1
	}

	//准备好排序的数据
	sortIndex := 0
	res = make([]int32, len(arr))
	for key, value := range countSlice {
		for i := int32(0); i < value; i++ {
			res[sortIndex] = int32(key) + minNum
			sortIndex++
		}
	}

	return
}
