package sort

//时间复杂度:平均:O(n+k)、最好:O(n^2)、最坏:O(n)
//空间复杂度:O(n+k)
//稳定性:稳定

/*
BucketSort
@Desc:桶排序
@Param:arr 			[]float64	待排数据
@Param:bucketCount 	int32		自定义桶个数
@Return:			[]float64	排序结果
*/
func BucketSort(arr []float64, bucketCount int32) []float64 {
	//定义桶的个数
	maxNum, _ := getMaxNumFloat64(arr)
	if bucketCount <= 0 {
		bucketCount = int32(len(arr))
	}

	span := maxNum / float64(bucketCount-1)     //桶区间跨度 = 最大值 / 桶个数-1
	bucketArr := make([][]float64, bucketCount) //生成桶

	//将数值分配到各个桶内
	for _, value := range arr {
		bucketKey := int32(value / span)
		bucketArr[bucketKey] = append(bucketArr[bucketKey], value)
	}

	//桶内排序
	sortRes := make([]float64, 0)
	for _, value := range bucketArr {
		res := InsertSortFloat64(value)
		sortRes = append(sortRes, res...)
	}

	return sortRes
}
