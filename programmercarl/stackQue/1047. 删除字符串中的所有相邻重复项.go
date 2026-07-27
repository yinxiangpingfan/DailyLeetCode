package stackque

func removeDuplicates(s string) string {
	ss := []byte(s)
	a := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if i == 0 {
			a = append(a, ss[i])
			continue
		}
		if len(a) != 0 && ss[i] == a[len(a)-1] {
			a = a[:len(a)-1]
		} else {
			a = append(a, ss[i])
		}

	}
	return string(a)
}
