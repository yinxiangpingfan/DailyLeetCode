package hot100

func maxSubArray(nums []int) int {
	a := make([]int, 0)
	maxs := -10086
	a = append(a, nums[0])
	if a[0] > maxs {
		maxs = a[0]
	}
	for i := 1; i < len(nums); i++ {
		a = append(a, max(nums[i], a[i-1]+nums[i]))
		if a[i] > maxs {
			maxs = a[i]
		}
	}
	return maxs
}
