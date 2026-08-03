package hot100

func subarraySum(nums []int, k int) int {
	count := 0
	n := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			n[i] = nums[i]
			continue
		}
		n[i] = n[i-1] + nums[i]
	}
	a := make(map[int]int)
	for j, _ := range n {
		v := n[len(nums)-1-j]
		count += a[v]
		a[v-k]++
	}
	count += a[0]
	return count
}
