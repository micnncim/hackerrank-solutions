package abbreviation

import (
	"fmt"
	"strings"
)

func abbreviation(a string, b string) bool {
	a = strings.ToUpper(a)
	for i := range a {
		if string(a[i]) == string(b[0]) {
			if len(b) == 1 {
				return true
			}
			if i == len(a) {
				return false
			}
			fmt.Printf("match %s a=%6s b=%6s\n", string(a[i]), a, b)
			abbreviation(a[i+1:], b[1:])
		} else {
			if i == len(a) {
				return false
			}
			fmt.Printf("unmatch a=%6s b=%6s\n", a, b)
			abbreviation(a[i+1:], b)
		}
	}
	return false
}
