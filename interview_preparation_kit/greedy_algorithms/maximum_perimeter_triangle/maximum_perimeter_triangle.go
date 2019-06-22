package maximum_perimeter_triangle

import (
	"sort"
)

func maximumPerimeterTriangle(sticks []int32) []int32 {
	sort.Slice(sticks, func(i, j int) bool {
		return sticks[i] > sticks[j]
	})
	for i := 0; i < len(sticks); i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < j; k++ {
				a, b, c := sticks[i], sticks[j], sticks[k]
				if a+b > c && b+c > a && c+a > b {
					s := []int32{a, b, c}
					sort.Slice(s, func(i, j int) bool {
						return s[i] < s[j]
					})
					return s
				}
			}
		}
	}
	return []int32{-1}
}
