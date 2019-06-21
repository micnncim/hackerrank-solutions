package arrays_left_rotation

func rotLeft(a []int32, d int32) []int32 {
	ret := make([]int32, 0, len(a))
	for i := range a {
		ret = append(ret, a[(int32(i)+d)%int32(len(a))])
	}
	return ret
}
