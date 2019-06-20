package array

func hourglassSum(arr [][]int32) int32 {
	max := int32(-63)
	for i := 0; i < len(arr)-2; i++ {
		for j := 0; j < len(arr[i])-2; j++ {
			p := arr[i][j] + arr[i][j+1] + arr[i][j+2] +
				arr[i+1][j+1] +
				arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			println(i, ",", j, ":", p)
			if p > max {
				max = p
			}
		}
	}
	return max
}
