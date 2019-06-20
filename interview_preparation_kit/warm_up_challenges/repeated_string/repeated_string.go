package repeated_string

func repeatedString(s string, n int64) int64 {
	var base int64
	var sub int64
	c := 'a'
	r := n % int64(len(s))
	for i := range s {
		if s[i] != uint8(c) {
			continue
		}
		if i < int(r) {
			sub++
		}
		base++
	}
	m := n / int64(len(s))
	return (base * m) + sub
}
