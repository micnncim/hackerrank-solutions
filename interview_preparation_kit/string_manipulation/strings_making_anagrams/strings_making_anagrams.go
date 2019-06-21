package strings_making_anagrams

func makeAnagram(a string, b string) int32 {
	var ret int32 = 0
	ma := make(map[string]int)
	for _, c := range a {
		ma[string(c)]++
	}
	for _, c := range b {
		if cnt, ok := ma[string(c)]; !ok || cnt <= 0 {
			ret++
		}
		ma[string(c)]--
	}
	mb := make(map[string]int)
	for _, c := range b {
		mb[string(c)]++
	}
	for _, c := range a {
		if cnt, ok := mb[string(c)]; !ok || cnt <= 0 {
			ret++
		}
		mb[string(c)]--
	}
	return ret
}
