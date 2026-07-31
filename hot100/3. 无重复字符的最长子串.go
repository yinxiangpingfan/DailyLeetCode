package hot100

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	if len(s) == 0 {
		return 0
	}
	maxCount := -1
	b := []byte(s)
	used := make([]int, 128)
	leftP := 0
	rightP := 1
	used[b[leftP]]++
	for rightP < len(b) {
		if used[b[rightP]] > 0 {
			//之前存在过
			used[b[leftP]]--
			leftP++
			if rightP-leftP+1 > maxCount {
				maxCount = rightP - leftP
			}
		} else {
			//之前没有存在过
			used[b[rightP]]++
			rightP++
			if rightP-leftP+1 > maxCount {
				maxCount = rightP - leftP
			}
		}
	}
	return maxCount
}
