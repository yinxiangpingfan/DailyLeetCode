package nums

func sortedSquares(nums []int) []int {
	leftP := 0
	rightP := len(nums) - 1
	slices := make([]int, rightP+1)
	count := rightP
	for leftP <= rightP {
		leftPP := nums[leftP] * nums[leftP]
		rightPP := nums[rightP] * nums[rightP]
		if leftPP < rightPP {
			slices[count] = rightPP
			count--
			rightP--
		} else {
			leftP++
			slices[count] = leftPP
			count--
		}
	}
	return slices
}
