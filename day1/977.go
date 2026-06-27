package day1

func sortedSquares(nums []int) []int {
	leftP := 0
	returnNums := make([]int, len(nums))
	countP := len(returnNums) - 1
	rightP := len(nums) - 1
	for leftP <= rightP {
		leftPN := nums[leftP] * nums[leftP]
		rightpN := nums[rightP] * nums[rightP]
		if leftPN <= rightpN {
			returnNums[countP] = rightpN
			rightP--
			countP--
		} else {
			returnNums[countP] = leftPN
			countP--
			leftP++
		}
	}
	return returnNums
}
