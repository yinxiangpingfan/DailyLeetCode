package doubleptr

func removeElement(nums []int, val int) int {
	firstp := 0
	rightp := 0
	for rightp <= len(nums)-1 {
		if nums[rightp] == val {
			rightp++
			continue
		}
		nums[firstp] = nums[rightp]
		firstp++
		rightp++
	}
	return firstp
}
