package max_array_sum

func maxSubsetSum(arr []int32) int32 {
	switch len(arr) {
	case 0:
		return 0
	case 1:
		return arr[0]
	case 2:
		return max(arr[0], arr[1])
	}
	tracker := make([]int32, len(arr))
	tracker[0] = arr[0]
	tracker[1] = max(arr[0], arr[1])
	for i := 2; i < len(arr); i++ {
		tracker[i] = max(tracker[i-1], arr[i], tracker[i-2]+arr[i])
	}
	return tracker[len(tracker)-1]
}

func max(values ...int32) int32 {
	var m int32
	for _, v := range values {
		if v > m {
			m = v
		}
	}
	return m
}
