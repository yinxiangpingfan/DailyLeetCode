package strings

func repeatedSubstringPattern(s string) bool {
	s1 := s + s
	s1 = s1[1 : len(s1)-1]
	sa := []byte(s)
	s1a := []byte(s1)
	//kmp
	next := make([]int, len(s))
	for i := 1; i < len(next); i++ {
		next[i] = next[i-1]

		for next[i] > 0 && sa[i] != sa[next[i]] {
			next[i] = next[next[i]-1]
		}

		if sa[i] == sa[next[i]] {
			next[i]++
		}
	}

	cur := 0
	for i := 0; i < len(s1a); i++ {

		for cur > 0 && sa[cur] != s1a[i] {
			cur = next[cur-1]
		}

		if sa[cur] == s1a[i] {
			cur++
		}

		if cur == len(sa) {
			return true
		}
	}
	return false
}
