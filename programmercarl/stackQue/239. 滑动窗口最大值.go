package stackque

func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0)
	rightP := 0 + k - 1
	leftP := 0
	rs := make([]int, 0)
	for {
		if leftP == 0 {
			queue = append(queue, leftP)
			leftP++
			for leftP <= rightP {
				if len(queue)-1 >= 0 && nums[leftP] >= nums[queue[len(queue)-1]] {
					queue = queue[:len(queue)-1]
				} else {
					queue = append(queue, leftP)
					leftP++
				}
			}
			rs = append(rs, nums[queue[0]])
			rightP++
		} else {
			if rightP >= len(nums) {
				break
			}
			leftP = rightP - k + 1
			if len(queue)-1 >= 0 && nums[rightP] >= nums[queue[len(queue)-1]] {
				queue = queue[:len(queue)-1]
			} else {
				queue = append(queue, rightP)
				if queue[0] < leftP {
					queue = queue[1:]
				}
				rightP++
				rs = append(rs, nums[queue[0]])
			}
		}
	}
	return rs
}
