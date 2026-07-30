package hot100

func moveZeroes(nums []int) {
	leftP := 0
	rightP := 0
	for rightP <= len(nums)-1 {
		if nums[rightP] == 0 {
			rightP++
			continue
		}
		if nums[leftP] == 0 {
			nums[leftP], nums[rightP] = nums[rightP], nums[leftP]
			leftP++
			continue
		}
		leftP++
		rightP++
	}
}
