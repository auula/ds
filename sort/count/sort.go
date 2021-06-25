package count

func Sort(arr []int) {
	copy(arr, sorting(arr))
}

func sorting(arr []int) []int {
	// 计算出数列的最大长度
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	// 统计出现的次数
	countArrays := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		countArrays[arr[i]]++
	}
	sortedArrays := make([]int, len(arr))
	// 排序
	index := 0
	for i := 0; i < len(countArrays); i++ {
		for j := 0; j < countArrays[i]; j++ {
			sortedArrays[index] = i
			index += 1
		}
	}
	return sortedArrays
}
