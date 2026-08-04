package backtracking

import "strconv"

func letterCombinations(digits string) []string {
	board := make([][]byte, 10)
	board[0] = []byte{
		' ',
	}
	board[1] = []byte{
		' ',
	}
	cur := 2
	t := make([]byte, 0)
	for i := 0; i < 25; {
		t = append(t, byte('a'+i))
		if (i+1)%3 == 0 {
			if i == 17 {
				t = append(t, 's')
				i++
			}
			if i == 20 {
				t = append(t, 'v')
				i++
			}
			if i == 23 {
				t = append(t, 'y', 'z')
				i++
			}
			a := make([]byte, len(t))
			copy(a, t)
			board[cur] = a
			cur++
			t = make([]byte, 0)
		}
		i++
	}
	s := []byte(digits)
	path := make([]byte, 0)
	rs := make([]string, 0)
	var backTrcaking func(s []byte, cur int)
	backTrcaking = func(s []byte, cur int) {
		if len(path) == len(digits) {
			rs = append(rs, string(path))
			return
		}
		e, _ := strconv.Atoi(string(s[cur]))
		for _, v := range board[e] {
			path = append(path, v)
			backTrcaking(s, cur+1)
			path = path[:len(path)-1]
		}
	}
	backTrcaking(s, 0)
	return rs
}
