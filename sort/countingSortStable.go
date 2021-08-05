package sort

//时间复杂度:平均:O(n+k)、最好:O(n+k)、最坏:O(n+k)
//空间复杂度:O(n+k)
//稳定性:不稳定

/*
CountingSortStable
@Desc:计数排序 - 稳定排序版本
@Param:arr 	[]CountingSortModel		待排数据
@Return:res []CountingSortModel		排序结果
*/
func CountingSortStable(arr []CountingSortModel) (res []CountingSortModel) {
	//统计次数
	maxNum, minNum := getMaxNumModel(arr)
	lenNum := maxNum - minNum + 1
	countSlice := make([]int32, lenNum)
	for _, value := range arr {
		countSlice[value.Sore-minNum] += 1
	}

	//数组变形统计
	sum := int32(0)
	for key, value := range countSlice {
		sum += value
		countSlice[key] = sum
	}

	//倒序遍历数据
	res = make([]CountingSortModel, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		countSliceKey := arr[i].Sore - minNum
		countSliceValue := countSlice[countSliceKey]
		res[countSliceValue-1] = arr[i]
		countSlice[countSliceKey]--
	}

	return
}
