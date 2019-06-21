package two_strings

func twoStrings(s1 string, s2 string) bool {
	m := make(map[string]bool)
	for _, c := range s1 {
		m[string(c)] = true
	}
	for _, c := range s2 {
		if _, ok := m[string(c)]; ok {
			return true
		}
	}
	return false
}
