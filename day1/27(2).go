package day1

//双指针的写法

func removeElement1(nums []int, val int) int {
	leftP := 0
	rightP := 0
	for rightP < len(nums) {
		if nums[rightP] == val {
			rightP++
		} else {
			nums[leftP] = nums[rightP]
			leftP++
			rightP++
		}
	}
	return leftP
}
