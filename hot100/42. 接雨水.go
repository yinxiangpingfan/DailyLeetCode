package hot100

func trap(height []int) int {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	rightMax[len(height)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		rightMax[len(height)-1-i] = max(rightMax[len(height)-i], height[len(height)-1-i])
	}
	count := 0
	for i := 0; i < len(height); i++ {
		count += (min(leftMax[i], rightMax[i]) - height[i])
	}
	return count
}
