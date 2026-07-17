package strings

import (
	"bufio"
	"fmt"
	"os"
)

func rightRotateString() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	nums := 0
	s := ""
	n, _ := fmt.Fscan(in, &nums, &s)
	if n == 0 {
		return
	}
	ss := []byte(s)
	l := 0
	r := len(ss) - 1
	for l <= r {
		t := ss[l]
		ss[l] = ss[r]
		ss[r] = t
		l++
		r--
	}
	l = 0
	r = nums - 1
	for l <= r {
		t := ss[l]
		ss[l] = ss[r]
		ss[r] = t
		l++
		r--
	}
	l = nums
	r = len(ss) - 1
	for l <= r {
		t := ss[l]
		ss[l] = ss[r]
		ss[r] = t
		l++
		r--
	}
	fmt.Fprintln(out, string(ss))
}
