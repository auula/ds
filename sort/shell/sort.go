package shell

import "github.com/higker/ds/sort"

type moreThan struct{}

func New() sort.Sorting {
	return &moreThan{}
}

func (m *moreThan) Sort(arr []float64) {
	pervIndex, current := 0, 0.0
	// 每次步长都是折半缩小增量
	for gap := len(arr); gap > 0; gap = gap / 2 {
		//	每次都是在特定步长进行比较
		for i := gap; i < len(arr); i++ {
			pervIndex = i - gap
			current = arr[i]
			for pervIndex >= 0 && arr[pervIndex] > current {
				arr[pervIndex+gap] = arr[pervIndex]
				pervIndex -= gap
			}
			arr[pervIndex+gap] = current
		}
	}
}
