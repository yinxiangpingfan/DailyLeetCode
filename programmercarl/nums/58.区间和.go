package nums

import (
	"bufio"
	"fmt"
	"os"
)

func rangeSum() {
	row := 0
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	n, _ := fmt.Fscan(in, &row)
	if n == 0 {
		return
	}
	a := make([]int, row)
	for i := 0; i < row; i++ {
		b := 0
		n, _ := fmt.Fscan(in, &b)
		if n == 0 {
			break
		}
		a[i] = b
	}
	//计算前缀和
	p := make([]int, row)
	for i := 0; i < row; i++ {
		if i == 0 {
			p[0] = a[0]
			continue
		}
		p[i] = p[i-1] + a[i]
	}
	left, right := 0, 0
	for {
		n, _ := fmt.Fscan(in, &left, &right)
		if n == 0 {
			break
		}
		nums := 0
		if left == 0 {
			nums = p[right]
		} else {
			nums = p[right] - p[left-1]
		}
		fmt.Fprintln(out, nums)
	}
}
