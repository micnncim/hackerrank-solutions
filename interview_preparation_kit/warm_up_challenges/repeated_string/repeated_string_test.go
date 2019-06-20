package repeated_string

import "testing"

func Test_repeatedString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		n    int64
		want int64
	}{
		{
			s:    "aba",
			n:    10,
			want: 7,
		},
		{
			s:    "a",
			n:    1000000000000,
			want: 1000000000000,
		},
		{
			s:    "ceebbcb",
			n:    817723,
			want: 0,
		},
		{
			s:    "gfcaaaecbg",
			n:    547602,
			want: 164280,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedString(tt.s, tt.n); got != tt.want {
				t.Errorf("repeatedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
