package roads_and_libraries

import (
	"testing"
)

func Test_roadsAndLibraries(t *testing.T) {
	tests := []struct {
		name   string
		n      int32
		c_lib  int32
		c_road int32
		cities [][]int32
		want   int64
	}{
		{
			name:   "",
			n:      3,
			c_lib:  2,
			c_road: 1,
			cities: [][]int32{
				{1, 2},
				{3, 1},
				{2, 3},
			},
			want: 4,
		},
		{
			name:   "",
			n:      6,
			c_lib:  2,
			c_road: 5,
			cities: [][]int32{
				{1, 3},
				{3, 4},
				{2, 4},
				{1, 2},
				{2, 3},
				{5, 6},
			},
			want: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roadsAndLibraries(tt.n, tt.c_lib, tt.c_road, tt.cities); got != tt.want {
				t.Errorf("roadsAndLibraries() = %v, want %v", got, tt.want)
			}
		})
	}
}
