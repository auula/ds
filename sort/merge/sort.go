package merge

type moreThan struct{}

func (m *moreThan) Sort(arr []float64) {
	mergeSort(arr)
}

func mergeSort(arr []float64) []float64 {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[:middle]
	right := arr[middle:]
	return func(left, right []float64) []float64 {
		var result []float64

		for len(left) != 0 && len(right) != 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		}

		if len(left) != 0 {
			result = append(result, left...)
		}
		if len(right) != 0 {
			result = append(result, right...)
		}

		return result
	}(mergeSort(left), mergeSort(right))
}
