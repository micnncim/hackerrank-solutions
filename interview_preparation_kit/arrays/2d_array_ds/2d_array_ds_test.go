package array

import "testing"

func Test_hourglassSum(t *testing.T) {
	tests := []struct {
		name string
		arr  [][]int32
		want int32
	}{
		{
			arr: [][]int32{
				{1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 0},
				{0, 0, 2, 4, 4, 0},
				{0, 0, 0, 2, 0, 0},
				{0, 0, 1, 2, 4, 0},
			},
			want: 19,
		},
		{
			arr: [][]int32{
				{1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 0},
				{0, 9, 2, -4, -4, 0},
				{0, 0, 0, -2, 0, 0},
				{0, 0, -1, -2, -4, 0},
			},
			want: 13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hourglassSum(tt.arr); got != tt.want {
				t.Errorf("hourglassSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
