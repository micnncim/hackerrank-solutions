package new_year_chaos

import (
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout

func minimumBribes(q []int32) {
	var result int32
	for i := len(q) - 1; i >= 0; i-- {
		if (q[i]-1)-int32(i) > 2 {
			fmt.Fprintf(stdout, "Too chaotic")
			return
		}
		for j := max(0, q[i]-2); j < int32(i); j++ {
			if q[j] > q[i] {
				result++
			}
		}
	}
	fmt.Fprint(stdout, result)
}

func max(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}
