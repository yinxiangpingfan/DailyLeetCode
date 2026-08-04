package backtracking

func combinationSum3(k int, n int) [][]int {
	count := 0
	path := make([]int, 0)
	ri := make([][]int, 0)
	var backtracking func(s int)
	backtracking = func(s int) {
		if len(path) == k {
			if count == n {
				a := make([]int, k)
				copy(a, path)
				ri = append(ri, a)
			}
		}
		for i := s; i < 10; i++ {
			if count+i <= n && len(path)+1 <= k {
				count += i
				path = append(path, i)
				backtracking(i + 1)
				count -= i
				path = path[:len(path)-1]
			}
		}
	}
	backtracking(1)
	return ri
}
