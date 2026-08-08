package backtracking

func findSubsequences(nums []int) [][]int {
	path := make([]int, 0)
	ri := make([][]int, 0)
	var backTraversol func(cur int)
	backTraversol = func(cur int) {
		if len(path) > 1 {
			t := make([]int, len(path))
			copy(t, path)
			ri = append(ri, t)
		}
		if cur == len(nums) {
			return
		}
		used := make(map[int]int)
		for i := cur; i < len(nums); i++ {
			if i > cur && nums[i] == nums[i-1] {
				continue
			}
			if (len(path) > 0 && nums[i] >= path[len(path)-1]) || len(path) == 0 {
				if used[nums[i]] == 0 {
					path = append(path, nums[i])
					used[nums[i]]++
					backTraversol(i + 1)
					path = path[:len(path)-1]
				} else {
					continue
				}
			}
		}
	}
	backTraversol(0)
	return ri
}
