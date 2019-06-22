package luck_balance

import (
	"sort"
)

func luckBalance(k int32, contests [][]int32) int32 {
	var ret int32
	important := make([]int32, 0, len(contests))
	for _, c := range contests {
		if c[1] == 1 {
			important = append(important, c[0])
		} else {
			ret += c[0]
		}
	}
	sort.Slice(important, func(i, j int) bool {
		return important[i] > important[j]
	})
	if len(important) == 0 {
		return ret
	}
	n := int32(len(important))
	if len(important) > int(k) {
		n = k
	}
	for _, v := range important[:n] {
		ret += v
	}
	for _, v := range important[n:] {
		ret -= v
	}
	return ret
}
