package math

import "sort"

func CalcAvgAndMedian(list []float64) (avg float64, median float64) {
	total := 0.0
	for _, item := range list {
		total += item
	}
	if len(list) == 0 {
		return 0, 0
	}
	avg = total / float64(len(list))

	sortedList := make([]float64, len(list))
	copy(sortedList, list)
	sort.Float64s(sortedList)

	mid := len(sortedList) / 2
	if len(sortedList)%2 == 0 {
		median = (sortedList[mid-1] + sortedList[mid]) / 2
	} else {
		median = sortedList[mid]
	}
	return avg, median
}
