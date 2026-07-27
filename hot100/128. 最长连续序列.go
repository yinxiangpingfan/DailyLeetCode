package hot100

func longestConsecutive(nums []int) int {
	a := make(map[int]struct{})
	for _, v := range nums {
		a[v] = struct{}{}
	}
	maxCount := 0
	for v, _ := range a {
		if _, ok := a[v-1]; ok {
			continue
		}
		count := 1
		for i := 1; ; i++ {
			if _, ok := a[i+v]; ok {
				count++
			} else {
				break
			}
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}
