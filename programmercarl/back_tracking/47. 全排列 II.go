package backtracking

func permuteUnique(nums []int) [][]int {
	ri := make([][]int, 0)
	path := make([]int, 0)
	used := make(map[int]int)
	for _, v := range nums {
		used[v]++
	}
	var backTracking func()
	backTracking = func() {
		if len(path) == len(nums) {
			t := make([]int, len(path))
			copy(t, path)
			ri = append(ri, t)
			return
		}
		uses := make(map[int]int)
		for i := 0; i < len(nums); i++ {
			if uses[nums[i]] > 0 {
				continue
			}
			if used[nums[i]] != 0 {
				path = append(path, nums[i])
				uses[nums[i]]++
				used[nums[i]]--
				backTracking()
				used[nums[i]]++
				path = path[:len(path)-1]
			}
		}
	}
	backTracking()
	return ri
}
