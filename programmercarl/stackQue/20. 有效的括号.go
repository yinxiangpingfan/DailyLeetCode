package stackque

func isValid(s string) bool {
	a := make([]byte, 0)
	b := make(map[byte]byte)
	b['{'] = '}'
	b['('] = ')'
	b['['] = ']'
	s1 := []byte(s)
	for i := 0; i < len(s1); i++ {
		if len(a) != 0 && s1[i] == b[a[len(a)-1]] {
			a = a[:len(a)-1]
		} else {
			a = append(a, s1[i])
		}
	}
	if len(a) == 0 {
		return true
	} else {
		return false
	}
}
