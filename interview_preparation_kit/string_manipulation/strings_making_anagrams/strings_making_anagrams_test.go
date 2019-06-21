package strings_making_anagrams

import "testing"

func Test_makeAnagram(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want int32
	}{
		{
			a:    "cde",
			b:    "abc",
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeAnagram(tt.a, tt.b); got != tt.want {
				t.Errorf("makeAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
