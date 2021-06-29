package bucket

import "github.com/higker/ds/sort/merge"

var (
	buckets [][]float64
	mg      = merge.New()
)

func Sort(arr []float64, size int) {

	max, min := arr[0], arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	buckets = make([][]float64, int(max-min)/(size)+1)

	for i := 0; i < len(arr); i++ {
		buckets[int(arr[i]-min)/size] = append(buckets[int(arr[i]-min)/size], arr[i])
	}
	index := 0
	for _, b := range buckets {
		if len(b) <= 0 {
			continue
		}
		mg.Sort(b)
		for _, v := range b {
			arr[index] = v
			index++
		}
	}
}
