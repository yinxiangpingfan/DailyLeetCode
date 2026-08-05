package backtracking

import "strconv"

func restoreIpAddresses(s string) []string {
	path := make([]byte, 0)
	rs := make([]string, 0)
	ss := []byte(s)
	count := 0
	var backTracking func(cur int)
	backTracking = func(cur int) {
		if count == 4 {
			if cur == len(s) {
				rs = append(rs, string(path[:len(path)-1]))
			}
			return
		}
		for i := cur; i < len(s); i++ {
			v := ss[cur : i+1]
			if len(v) > 1 && v[0] == '0' {
				break
			}
			for _, v1 := range v {
				s := v1 - '0'
				if s > 9 {
					break
				}
			}
			a, _ := strconv.Atoi(string(v))
			if a < 0 || a > 255 {
				break
			}
			path = append(path, v...)
			path = append(path, '.')
			count++
			backTracking(i + 1)
			count--
			path = path[:len(path)-len(v)-1]
		}
	}
	backTracking(0)
	return rs
}
