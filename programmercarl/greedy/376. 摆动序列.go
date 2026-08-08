package greedy

func wiggleMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	t := 1
	cur := 2 // 0上 1下
	for i := 1; i < len(nums); {
		if cur == 2 {
			if i >= len(nums) {
				break
			}
			if nums[i] > nums[i-1] {
				t++
				cur = 0
			} else if nums[i] < nums[i-1] {
				t++
				cur = 1
			}
			i++
			continue
		}
		if nums[i] < nums[i-1] {
			if cur == 1 {
				//前一是下
			} else {
				cur = 1
				t++
			}
			i++
		} else if nums[i] > nums[i-1] {
			if cur == 0 {
				//前一是上
			} else {
				cur = 0
				t++
			}
			i++
		} else {
			i++
		}
	}
	return t
}
