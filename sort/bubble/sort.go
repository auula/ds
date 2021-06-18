package bubble

import "github.com/higker/ds/sort"

type moreThan struct{}
type lessThan struct{}

func (m *moreThan) Sort(arr []float64) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func (le *lessThan) Sort(arr []float64) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func New(symbol rune) sort.Sorting {
	switch symbol {
	case '<':
		return &lessThan{}
	default:
		return &moreThan{}
	}
}
