package doubleptr

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
	s1 := []byte(s)
	s2 := make([]byte, 0)
	for i := 0; i < len(s1); i++ {
		if s1[i] >= 'a' && s1[i] <= 'z' {
			s2 = append(s2, s1[i])
		} else {
			t := []byte{
				'n', 'u', 'm', 'b', 'e', 'r',
			}
			s2 = append(s2, t...)
		}
	}
	fmt.Fprintln(out, string(s2))
}
