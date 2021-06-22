package insertion

import "github.com/higker/ds/sort"

type moreThan struct{}

func New() sort.Sorting {
	return &moreThan{}
}

func (m *moreThan) Sort(arr []float64) {
	for i := range arr {
		// 拿到排序和没有排序的
		pervIndex := i - 1
		current := arr[i]
		// 把没有排序的和前面的比较
		for pervIndex >= 0 && arr[pervIndex] > current {
			arr[pervIndex+1] = arr[pervIndex]
			pervIndex -= 1
		}
		arr[pervIndex+1] = current
	}
}
