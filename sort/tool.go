package sort

// CountingSortModel 稳定计数排序测试用例数值
type CountingSortModel struct {
	Name string
	Sore int32
}


/*
getMaxNumFloat64
@Desc:获得数组内最大值和最小值
@Param:arr 	[]int32	数组
@Return:max int32	最大值
@Return:min int32	最小值
*/
func getMaxNumFloat64(arr []float64) (max float64, min float64) {
	if len(arr) == 0 {
		return
	}

	max = arr[0]
	min = arr[0]

	for _, value := range arr {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	return
}



/*
getMaxNumModel
@Desc:获得数组内最大值和最小值
@Param:arr 	[]int32	数组
@Return:max int32	最大值
@Return:min int32	最小值
*/
func getMaxNumModel(arr []CountingSortModel) (max int32, min int32) {
	if len(arr) == 0 {
		return
	}

	max = arr[0].Sore
	min = arr[0].Sore

	for _, value := range arr {
		if value.Sore > max {
			max = value.Sore
		}
		if value.Sore < min {
			min = value.Sore
		}
	}

	return
}

/*
getMaxNum
@Desc:获得数组内最大值和最小值
@Param:arr 	[]int32	数组
@Return:max int32	最大值
@Return:min int32	最小值
*/
func getMaxNum(arr []int32) (max int32, min int32) {
	if len(arr) == 0 {
		return
	}

	max = arr[0]
	min = arr[0]

	for _, value := range arr {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	return
}

