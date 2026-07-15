package nums

func removeElement(nums []int, val int) int {
	firstP := 0
	slowP := 0
	for slowP <= len(nums)-1 {
		if nums[slowP] != val {
			nums[firstP] = nums[slowP]
			firstP++
		}
		slowP++
	}
	return firstP
}
