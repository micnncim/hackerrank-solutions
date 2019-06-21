package minimum_absolute_difference_in_an_array

import (
	"math"
	"sort"
)

func minimumAbsoluteDifference(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	var min int32 = math.MaxInt32
	for i := 0; i < len(arr)-1; i++ {
		d := diff(arr[i], arr[i+1])
		if d < min {
			min = d
		}
	}
	return min
}

func diff(x, y int32) int32 {
	d := x - y
	if d > 0 {
		return d
	}
	return -d
}
