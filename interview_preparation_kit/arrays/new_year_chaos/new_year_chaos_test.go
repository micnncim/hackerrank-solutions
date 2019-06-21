package new_year_chaos

import (
	"bytes"
	"testing"
)

func Test_minimumBribes(t *testing.T) {
	tests := []struct {
		name string
		q    []int32
		want string
	}{
		{
			q:    []int32{2, 1, 5, 3, 4},
			want: "3",
		},
		{
			q:    []int32{1, 2, 5, 3, 7, 8, 6, 4},
			want: "7",
		},
		{
			q:    []int32{2, 5, 1, 3, 4},
			want: "Too chaotic",
		},
		{
			q:    []int32{5, 1, 2, 3, 7, 8, 6, 4},
			want: "Too chaotic",
		},
	}

	for _, tt := range tests {
		buf := new(bytes.Buffer)
		stdout = buf

		t.Run(tt.name, func(t *testing.T) {
			minimumBribes(tt.q)
			if buf.String() != tt.want {
				t.Errorf("minimumBribes() = %v, want %v", buf.String(), tt.want)
			}
		})
	}
}
