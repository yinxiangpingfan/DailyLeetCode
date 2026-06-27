package day1

// 暴力解法
func removeElement(nums []int, val int) int {
	count := 0
	k := 0
	t := 0
	for t < len(nums) {
		if nums[k] == val {
			//移动
			for i := k + 1; i < len(nums); i++ {
				nums[i-1] = nums[i]
			}
		} else {
			count++
			k++
		}
		t++
	}
	return count
}
