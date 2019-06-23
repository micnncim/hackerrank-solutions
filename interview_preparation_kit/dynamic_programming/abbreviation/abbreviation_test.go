package abbreviation

import "testing"

func Test_abbreviation(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want bool
	}{
		{
			a:    "daBcd",
			b:    "ABC",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abbreviation(tt.a, tt.b); got != tt.want {
				t.Errorf("abbreviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
