package strings

import (
	"bufio"
	"fmt"
	"os"
)

func replaceNumber() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	s := ""
	n, _ := fmt.Fscan(in, &s)
	if n == 0 {
		return
	}
	b := []byte(s)
	c := make([]byte, 0)
	for _, v := range b {
		if v >= 'a' && v <= 'z' {
			c = append(c, v)
		} else {
			number := []byte{
				'n', 'u', 'm', 'b', 'e', 'r',
			}
			c = append(c, number...)
		}
	}
	fmt.Fprintln(out, string(c))
}
