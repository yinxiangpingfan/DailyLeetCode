package backtracking

func solveSudoku(board [][]byte) {
	var backTRaversol func(board [][]byte) bool
	backTRaversol = func(board [][]byte) bool {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {
				if board[i][j] == '.' {
					for v := 1; v <= 9; v++ {
						if isOk1(i, j, byte(v+'0'), board) {
							board[i][j] = byte(v + '0')
							if backTRaversol(board) {
								return true
							}
							board[i][j] = '.'
						}
					}
					return false
				} else {
					continue
				}
			}
		}
		return true
	}
	backTRaversol(board)
}

func isOk1(row, col int, k byte, board [][]byte) bool {
	for i := 0; i < 9; i++ { //行
		if board[row][i] == k {
			return false
		}
	}
	for i := 0; i < 9; i++ { //列
		if board[i][col] == k {
			return false
		}
	}
	//方格
	startrow := (row / 3) * 3
	startcol := (col / 3) * 3
	for i := startrow; i < startrow+3; i++ {
		for j := startcol; j < startcol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}
