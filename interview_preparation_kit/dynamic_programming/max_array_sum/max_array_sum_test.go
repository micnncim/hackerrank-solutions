package max_array_sum

import "testing"

func Test_maxSubsetSum(t *testing.T) {
	tests := []struct {
		name string
		arr  []int32
		want int32
	}{
		{
			arr:  []int32{3, 7, 4, 6, 5},
			want: 13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubsetSum(tt.arr); got != tt.want {
				t.Errorf("maxSubsetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
