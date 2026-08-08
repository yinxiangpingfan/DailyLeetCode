package greedy

func maxSubArray(nums []int) int {
	t := nums[0]
	max := t
	if len(nums) == 1 {
		return t
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]+t > nums[i] {
			t = nums[i] + t
		} else {
			t = nums[i]
		}
		if t > max {
			max = t
		}
	}
	return max
}
