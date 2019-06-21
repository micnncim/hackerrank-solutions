package alternating_characters

import "testing"

func Test_alternatingCharacters(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int32
	}{
		{s: "AAAA", want: 3},
		{s: "ABABABAB", want: 0},
		{s: "AAABBB", want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alternatingCharacters(tt.s); got != tt.want {
				t.Errorf("alternatingCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
