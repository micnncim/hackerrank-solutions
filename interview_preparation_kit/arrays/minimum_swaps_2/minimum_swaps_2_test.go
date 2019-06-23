package minimum_swaps_2

import "testing"

func Test_minimumSwaps(t *testing.T) {
	tests := []struct {
		name string
		arr  []int32
		want int32
	}{
		{
			arr:  []int32{4, 3, 1, 2},
			want: 3,
		},
		{
			arr:  []int32{2, 1, 4, 3},
			want: 2,
		},
		{
			arr:  []int32{2, 3, 4, 1, 5},
			want: 3,
		},
		{
			arr:  []int32{1, 3, 5, 2, 4, 6, 7},
			want: 3,
		},
		{
			arr:  []int32{4, 3, 1, 2},
			want: 3,
		},
		{
			arr:  []int32{3, 7, 6, 9, 1, 8, 10, 4, 2, 5},
			want: 9,
		},
		{
			arr:  []int32{2, 31, 1, 38, 29, 5, 44, 6, 12, 18, 39, 9, 48, 49, 13, 11, 7, 27, 14, 33, 50, 21, 46, 23, 15, 26, 8, 47, 40, 3, 32, 22, 34, 42, 16, 41, 24, 10, 4, 28, 36, 30, 37, 35, 20, 17, 45, 43, 25, 19},
			want: 46,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSwaps(tt.arr); got != tt.want {
				t.Errorf("minimumSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
