package strings

func strStr(haystack string, needle string) int {
	//kmp解法
	ha := []byte(haystack)
	ne := []byte(needle)
	//求next
	next := make([]int, len(needle))
	for i := 1; i < len(ne); i++ {
		next[i] = next[i-1]

		for next[i] > 0 && ne[i] != ne[next[i]] {
			next[i] = next[next[i]-1]
		}

		if ne[i] == ne[next[i]] {
			next[i]++
		}
	}

	if len(ha) < len(ne) {
		return -1
	}

	cur := 0
	for i := 0; i < len(ha); i++ {
		for cur > 0 && ha[i] != ne[cur] {
			cur = next[cur-1]
		}
		if ha[i] == ne[cur] {
			cur++
		}
		if cur == len(ne) {
			return i - len(ne) + 1
		}
	}
	return -1
}
