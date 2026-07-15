package nums

import (
	"bufio"
	"fmt"
	"os"
)

func buyRoil() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	hang, lie := 0, 0
	n, _ := fmt.Fscan(in, &hang, &lie)
	if n == 0 {
		return
	}
	square := make([][]int, hang)
	for i := 0; i < hang; i++ {
		square[i] = make([]int, lie)
	}
	for i := 0; i < hang*lie; i++ {
		a := 0
		n, _ := fmt.Fscan(in, &a)
		if n == 0 {
			break
		}
		square[i/lie][i%lie] = a
	}
	//二维前缀和
	pHang := make([]int, hang)
	pLie := make([]int, lie)
	for i := 0; i < hang; i++ {
		count := 0
		for j := 0; j < lie; j++ {
			count += square[i][j]
		}
		if i == 0 {
			pHang[i] = count
		} else {
			pHang[i] = count + pHang[i-1]
		}
	}

	for i := 0; i < lie; i++ {
		count := 0
		for j := 0; j < hang; j++ {
			count += square[j][i]
		}
		if i == 0 {
			pLie[i] = count
		} else {
			pLie[i] = count + pLie[i-1]
		}
	}
	min := 10001
	//按行分
	for i := 0; i < hang-1; i++ {
		leftPPPP := pHang[i]
		rightPPP := pHang[hang-1] - leftPPPP
		if leftPPPP > rightPPP {
			c := leftPPPP - rightPPP
			if c < min {
				min = c
			}
		} else {
			c := rightPPP - leftPPPP
			if c < min {
				min = c
			}
		}
	}
	//按列分
	for i := 0; i < lie-1; i++ {
		leftPPPP := pLie[i]
		rightPPP := pLie[lie-1] - leftPPPP
		if leftPPPP > rightPPP {
			c := leftPPPP - rightPPP
			if c < min {
				min = c
			}
		} else {
			c := rightPPP - leftPPPP
			if c < min {
				min = c
			}
		}
	}
	fmt.Fprintln(out, min)
}
