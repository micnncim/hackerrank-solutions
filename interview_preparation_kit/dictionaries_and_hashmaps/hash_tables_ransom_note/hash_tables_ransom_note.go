package hash_tables_ransom_note

func checkMagazine(magazine []string, note []string) bool {
	m := make(map[string]int)
	for _, word := range magazine {
		m[word]++
	}
	for _, word := range note {
		if c, ok := m[word]; !ok || c == 0 {
			return false
		}
		m[word]--
	}
	return true
}
