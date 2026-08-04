package backtracking

func combine2(n int, k int) [][]int {
	ri := make([][]int, 0)
	path := make([]int, 0)
	var backtracking func(s int)
	backtracking = func(s int) {
		if len(path) == k {
			a := make([]int, k)
			copy(a, path)
			ri = append(ri, a)
			return
		}
		for i := s; i <= n; i++ {
			if n-i+1 < k-len(path) {
				return
			}
			path = append(path, i)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(1)
	return ri
}
