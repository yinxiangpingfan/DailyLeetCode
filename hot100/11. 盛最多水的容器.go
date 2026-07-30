package hot100

func maxArea(height []int) int {
	leftP := 0
	rightP := len(height) - 1
	max := -1
	for leftP <= rightP {
		min := 0
		s := 0
		if height[leftP] < height[rightP] {
			min = height[leftP]
			s = (rightP - leftP) * min
			leftP++
		} else {
			min = height[rightP]
			s = (rightP - leftP) * min
			rightP--
		}
		if s > max {
			max = s
		}
	}
	return max
}

func mina(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
