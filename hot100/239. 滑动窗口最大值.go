package hot100

func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	rn := make([]int, 0)
	maxQ := make([]int, 1)
	maxQ[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		s := 0
		if i >= k {
			s = i - k
			if nums[s] == maxQ[0] {
				maxQ = maxQ[1:]
			}
		}
		if nums[i] <= maxQ[len(maxQ)-1] {
			maxQ = append(maxQ, nums[i])
		} else {
			for len(maxQ) > 0 && nums[i] > maxQ[len(maxQ)-1] {
				maxQ = maxQ[:len(maxQ)-1]
			}
			maxQ = append(maxQ, nums[i])
		}
		if i >= k-1 {
			rn = append(rn, maxQ[0])
		}
	}
	return rn
}
