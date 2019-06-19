package jumping_on_the_clouds

import "testing"

func Test_jumpingOnClouds(t *testing.T) {
	tests := []struct {
		name string
		c    []int32
		want int32
	}{
		{
			name: "",
			c:    []int32{0, 0, 1, 0, 0, 1, 0},
			want: 4,
		},
		{
			name: "",
			c:    []int32{0, 0, 0, 1, 0, 0},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jumpingOnClouds(tt.c); got != tt.want {
				t.Errorf("jumpingOnClouds() = %v, want %v", got, tt.want)
			}
		})
	}
}
