package backtracking

func subsets(nums []int) [][]int {
	path := make([]int, 0)
	ri := make([][]int, 0)
	var backTracking func(cur int)
	backTracking = func(cur int) {
		t := make([]int, len(path))
		copy(t, path)
		ri = append(ri, t)

		if cur == len(nums) {
			return
		}

		for i := cur; i < len(nums); i++ {
			path = append(path, nums[i])
			backTracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTracking(0)
	return ri
}
