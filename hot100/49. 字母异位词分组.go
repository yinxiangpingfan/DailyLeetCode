package hot100

import "sort"

func groupAnagrams(strs []string) [][]string {
	a := make(map[string][]string)
	for _, v := range strs {
		t := []byte(v)
		sort.Slice(t, func(i, j int) bool {
			return t[i] < t[j]
		})
		a[string(t)] = append(a[string(t)], v)
	}
	rs := make([][]string, 0)
	for _, v := range a {
		rs = append(rs, v)
	}
	return rs
}
