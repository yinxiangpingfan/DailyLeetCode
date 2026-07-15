package nums

func minSubArrayLen(target int, nums []int) int {
	leftP := 0
	rightP := 0
	max := 0
	count := 0
	returnInt := len(nums) + 1
	for rightP <= len(nums)-1 {
		max += nums[rightP]
		count++
		if max >= target {
			//大于了，开始缩短窗口
			if count < returnInt {
				returnInt = count
			}
			for leftP < rightP {
				max -= nums[leftP]
				count--
				leftP++
				if max < target {
					break
				} else {
					if count < returnInt {
						returnInt = count
					}
				}
			}
		}
		rightP++
	}
	if returnInt == len(nums)+1 {
		return 0
	} else {
		return returnInt
	}
}
