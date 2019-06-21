package arrays_left_rotation

import (
	"reflect"
	"testing"
)

func Test_rotLeft(t *testing.T) {
	tests := []struct {
		name string
		a    []int32
		d    int32
		want []int32
	}{
		{
			a:    []int32{1, 2, 3, 4, 5},
			d:    4,
			want: []int32{5, 1, 2, 3, 4},
		},
		{
			a:    []int32{41, 73, 89, 7, 10, 1, 59, 58, 84, 77, 77, 97, 58, 1, 86, 58, 26, 10, 86, 51},
			d:    10,
			want: []int32{77, 97, 58, 1, 86, 58, 26, 10, 86, 51, 41, 73, 89, 7, 10, 1, 59, 58, 84, 77},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotLeft(tt.a, tt.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}
