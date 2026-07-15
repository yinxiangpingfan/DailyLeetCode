package nums

func search(nums []int, target int) int {
	leftP := 0
	right := len(nums) - 1
	for leftP <= right {
		mid := (leftP + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			leftP = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
