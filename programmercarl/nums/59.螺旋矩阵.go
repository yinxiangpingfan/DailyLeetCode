package nums

func generateMatrix(n int) [][]int {
	//初始化
	square := make([][]int, n)
	for i := range len(square) {
		square[i] = make([]int, n)
	}
	//模拟
	// 1 2 3 4
	// 5 6 7 8
	// 9 10 11 12
	// 13 14 15 16
	if n%2 == 1 {
		square[n/2][n/2] = n * n
	}
	cur := n - 1
	top := 0
	left := 0
	cursor := 1
	for {
		if cur <= 0 {
			break
		}
		//上
		for i := 0; i < cur; i++ {
			square[top][i+left] = cursor
			cursor++
		}
		//右
		for i := 0; i < cur; i++ {
			square[i+top][n-1-left] = cursor
			cursor++
		}
		//下
		for i := 0; i < cur; i++ {
			square[n-1-top][n-1-i-left] = cursor
			cursor++
		}
		//左
		for i := 0; i < cur; i++ {
			square[n-1-top-i][left] = cursor
			cursor++
		}
		cur -= 2
		top++
		left++
	}
	return square
}
