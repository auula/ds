package count

// 量大但是范围小
func Sort(arr []int) {
	copy(arr, sorting(arr))
}

// 为不浪费空间 数组每个元素是某一个范围内 有点大材小用
// 下标计算就需要偏移量了min
func sorting(arr []int) []int {
	// 计算出数列的最大长度
	max, min := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	// 统计出现的次数 max-min+1
	different := max - min

	countArrays := make([]int, different+1)
	for i := 0; i < len(arr); i++ {
		// min 偏移量
		countArrays[arr[i]-min]++
	}

	// 整理变形 加上之前的前面下标的值
	for i := 1; i < len(arr); i++ {
		countArrays[i] += countArrays[i-1]
	}

	sortedArrays := make([]int, len(arr))
	// 排序
	for i := len(arr) - 1; i >= 0; i-- {
		sortedArrays[countArrays[arr[i]-min]-1] = arr[i]
		countArrays[arr[i]-min]--
	}
	return sortedArrays
}
