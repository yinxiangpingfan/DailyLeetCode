package stackque

import "sort"

func topKFrequent(nums []int, k int) []int {
	a := make(map[int]int)
	for _, v := range nums {
		a[v]++
	}
	b := make(map[int][]int)
	for k, v := range a {
		b[v] = append(b[v], k)
	}
	c := make([]int, 0)
	for k, _ := range b {
		c = append(c, k)
	}
	sort.Slice(c, func(i, j int) bool {
		return c[i] > c[j]
	})
	rs := make([]int, 0)
	for i := 0; i < k; i++ {
		if len(rs) >= k {
			break
		}
		t := b[c[i]]
		for j := 0; j < len(t); j++ {
			rs = append(rs, t[j])
			if len(rs) >= k {
				break
			}
		}
	}
	return rs
}
