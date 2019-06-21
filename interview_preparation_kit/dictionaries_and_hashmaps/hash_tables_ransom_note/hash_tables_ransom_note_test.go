package hash_tables_ransom_note

import "testing"

func Test_checkMagazine(t *testing.T) {
	tests := []struct {
		name     string
		magazine []string
		note     []string
		want     bool
	}{
		{
			magazine: []string{"give", "me", "one", "grand", "today", "night"},
			note:     []string{"give", "one", "grand", "today"},
			want:     true,
		},
		{
			magazine: []string{"two", "times", "three", "is", "not", "four"},
			note:     []string{"two", "times", "two", "is", "four"},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkMagazine(tt.magazine, tt.note); got != tt.want {
				t.Errorf("checkMagazine() = %v, want %v", got, tt.want)
			}
		})
	}
}
