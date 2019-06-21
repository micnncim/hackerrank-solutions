package minimum_absolute_difference_in_an_array

import "testing"

func Test_minimumAbsoluteDifference(t *testing.T) {
	tests := []struct {
		name string
		arr  []int32
		want int32
	}{
		{
			arr:  []int32{3, -7, 0},
			want: 3,
		},
		{
			arr:  []int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumAbsoluteDifference(tt.arr); got != tt.want {
				t.Errorf("minimumAbsoluteDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
