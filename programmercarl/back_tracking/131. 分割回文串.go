package backtracking

func partition(s string) [][]string {
	ss := []byte(s)
	var backTracking func(cur int)
	path := make([]string, 0)
	ri := make([][]string, 0)
	backTracking = func(cur int) {
		if cur == len(ss) {
			t := make([]string, len(path))
			copy(t, path)
			ri = append(ri, t)
		}
		for i := cur; i < len(ss); i++ {
			v := ss[cur : i+1]
			if isT(string(v)) {
				path = append(path, string(v))
				backTracking(i + 1)
				path = path[:len(path)-1]
			} else {
				continue
			}
		}
	}
	backTracking(0)
	return ri
}

func isT(ss string) bool {
	s := []byte(ss)
	leftP := 0
	rightP := len(s) - 1
	for leftP <= rightP {
		if s[leftP] == s[rightP] {
			leftP++
			rightP--
		} else {
			return false
		}
	}
	return true
}
