package hash_tables_ice_cream_parlor

func whatFlavors(cost []int32, money int32) []int32 {
	m := make(map[int32]int32)

	for i, c := range cost {
		diff := money - c
		if v, ok := m[diff]; ok {
			return []int32{v, int32(i)}
		}
		m[c] = int32(i)
	}

	return nil
}
