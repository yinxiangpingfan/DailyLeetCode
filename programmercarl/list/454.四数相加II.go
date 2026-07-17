package list

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	a := make(map[int]int)
	lens := len(nums1)
	for i := 0; i < lens; i++ {
		for j := 0; j < lens; j++ {
			b := nums1[i] + nums2[j]
			a[0-b] += 1
		}
	}
	count := 0
	for i := 0; i < lens; i++ {
		for j := 0; j < lens; j++ {
			b := nums3[i] + nums4[j]
			count += a[b]
		}
	}
	return count
}
