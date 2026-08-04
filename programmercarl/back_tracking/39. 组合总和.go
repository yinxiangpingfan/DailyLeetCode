package backtracking

func combinationSum(candidates []int, target int) [][]int {
	ri := make([][]int, 0)
	path := make([]int, 0)
	count := 0
	var backTraversol func(cur int)
	backTraversol = func(cur int) {
		if count == target {
			a := make([]int, len(path))
			copy(a, path)
			ri = append(ri, a)
		}
		for i := cur; i < len(candidates); i++ {
			if count < target {
				count += candidates[i]
				path = append(path, candidates[i])
				backTraversol(i)
				count -= candidates[i]
				path = path[:len(path)-1]
			}
		}
	}
	backTraversol(0)
	return ri
}
