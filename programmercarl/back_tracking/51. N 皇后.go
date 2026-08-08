package backtracking

func solveNQueens(n int) [][]string {
	board := make([][]string, 0)
	used := make(map[int]int)
	var backTraversol func(cur int)
	path := make([]string, 0)
	backTraversol = func(cur int) {
		if len(path) == n {
			t := make([]string, len(path))
			copy(t, path)
			board = append(board, t)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] == 0 {
				a := make([]byte, 0)
				if isOk(path, cur, i) {
					used[i]++
					for j := 0; j < i; j++ {
						a = append(a, '.')
					}
					a = append(a, 'Q')
					for j := i + 1; j < n; j++ {
						a = append(a, '.')
					}
					path = append(path, string(a))
					backTraversol(cur + 1)
					used[i]--
					path = path[:len(path)-1]
				}
			}
		}
	}
	backTraversol(0)
	return board
}

func isOk(path []string, x, y int) bool {
	board := make([][]byte, len(path))
	for k, v := range path {
		t := []byte(v)
		board[k] = t
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return true
	}
	a, b := x, y
	for {
		if a == 0 || b == 0 {
			break
		}
		a--
		b--
	}
	c, d := x, y
	for {
		if c == 0 || d == len(board[0])-1 {
			break
		}
		c--
		d++
	}
	for {
		if a == len(board) || b == len(board[0]) {
			break
		}
		if board[a][b] == 'Q' {
			return false
		}
		a++
		b++
	}
	for {
		if c == len(board) || d == -1 {
			break
		}
		if board[c][d] == 'Q' {
			return false
		}
		c++
		d--
	}
	return true
}
