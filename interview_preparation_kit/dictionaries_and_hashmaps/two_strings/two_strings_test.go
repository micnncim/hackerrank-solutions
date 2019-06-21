package two_strings

import "testing"

func Test_twoStrings(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{
			s1:   "hello",
			s2:   "world",
			want: true,
		},
		{
			s1:   "hi",
			s2:   "world",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoStrings(tt.s1, tt.s2); got != tt.want {
				t.Errorf("twoStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
