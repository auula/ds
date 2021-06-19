package quick

import "github.com/higker/ds/sort"

type (
	moreThan struct{}
	lessThan struct{}
)

func New() sort.Sorting {
	return &moreThan{}
}

// 快速排序
// 右小停，左大停，然后交换左右
// 大循环结束 移动重合指针的元素
func partition(arr []float64, startIndex, endIndex int) int {
	left, right := startIndex, endIndex
	pivot := arr[startIndex]
	for left != right {
		for left < right && arr[right] >= pivot {
			right--
		}
		for left < right && arr[left] <= pivot {
			left++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[startIndex] = arr[left]
	arr[left] = pivot
	return left
}

func quickSort(arr []float64, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	// 第一次循环区分中线，然后左右递归
	pivotIndex := partition(arr, startIndex, endIndex)
	quickSort(arr, startIndex, pivotIndex-1)
	quickSort(arr, pivotIndex+1, endIndex)
}

func (m *moreThan) Sort(arr []float64) {
	quickSort(arr, 0, len(arr)-1)
}
