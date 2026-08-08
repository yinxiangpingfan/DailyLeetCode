package greedy

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	count := 0
	for i := 0; i < len(s); i++ {
		if count >= len(g) {
			break
		}
		if s[i] >= g[count] {
			count++
		}
	}
	return count
}
