package hash_tables_ice_cream_parlor

import (
	"reflect"
	"testing"
)

func Test_whatFlavors(t *testing.T) {
	tests := []struct {
		name  string
		cost  []int32
		money int32
		want  []int32
	}{
		{
			cost:  []int32{1, 4, 5, 3, 2},
			money: 4,
			want:  []int32{0, 3},
		},
		{
			cost:  []int32{2, 2, 4, 3},
			money: 4,
			want:  []int32{0, 1},
		},
		{
			cost:  []int32{1, 4, 5, 3, 2},
			money: 4,
			want:  []int32{0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := whatFlavors(tt.cost, tt.money); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("whatFlavors() = %v, want %v", got, tt.want)
			}
		})
	}
}
