package maximum_perimeter_triangle

import (
	"reflect"
	"testing"
)

func Test_maximumPerimeterTriangle(t *testing.T) {
	tests := []struct {
		name   string
		sticks []int32
		want   []int32
	}{
		{
			sticks: []int32{1, 1, 1, 3, 3},
			want:   []int32{1, 3, 3},
		},
		{
			sticks: []int32{1, 2, 3},
			want:   []int32{-1},
		},
		{
			sticks: []int32{1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000, 1000000000},
			want:   []int32{1000000000, 1000000000, 1000000000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumPerimeterTriangle(tt.sticks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumPerimeterTriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
