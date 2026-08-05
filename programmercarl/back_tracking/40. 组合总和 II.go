package backtracking

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	path := make([]int, 0)
	ri := make([][]int, 0)
	count := 0
	var backTracking func(cur int)
	backTracking = func(cur int) {
		if count == target {
			t := make([]int, len(path))
			copy(t, path)
			ri = append(ri, t)
		}
		for i := cur; i < len(candidates); i++ {
			if i > cur && candidates[i] == candidates[i-1] {
				continue
			}
			if count+candidates[i] <= target {
				count += candidates[i]
				path = append(path, candidates[i])
				backTracking(i + 1)
				count -= candidates[i]
				path = path[:len(path)-1]
			} else {
				break
			}
		}
	}
	backTracking(0)
	return ri
}
