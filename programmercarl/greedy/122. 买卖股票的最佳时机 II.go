package greedy

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}
	}
	return max
}
