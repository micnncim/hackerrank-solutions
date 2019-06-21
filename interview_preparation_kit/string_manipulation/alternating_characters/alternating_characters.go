package alternating_characters

func alternatingCharacters(s string) int32 {
	var cnt int32 = 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			cnt++
		}
	}
	return cnt
}
