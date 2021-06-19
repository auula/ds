package selection

import "github.com/higker/ds/sort"

// 假定一个元素和后面比较

func selection(arr []float64) {
	min := 0
	for i := 0; i < len(arr)-1; i++ {
		min = i
		// 从被选择的下标后面一个开始
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

type moreThan struct{}

func (m *moreThan) Sort(arr []float64) {
	selection(arr)
}

func New() sort.Sorting {
	return &moreThan{}
}
