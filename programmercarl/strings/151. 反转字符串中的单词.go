package strings

func reverseWords(s string) string {
	leftP := 0
	rightP := leftP
	ss := []byte(s)
	//反转字符串
	reverse(ss, 0, len(ss)-1)
	//先去除空格
	ss = remove(ss)
	//给单词换位置
	for rightP < len(ss) {
		if ss[rightP] == ' ' {
			reverse(ss, leftP, rightP-1)
			leftP = rightP + 1
			rightP = leftP
		} else {
			rightP++
		}
	}
	if ss[len(ss)-1] != ' ' {
		reverse(ss, leftP, len(ss)-1)
	} else {
		return string(ss[:len(ss)-1])
	}
	return string(ss)
}

func remove(a []byte) []byte {
	b := make([]byte, 0)
	for i := 0; i < len(a); i++ {
		if a[i] != ' ' {
			b = append(b, a[i])
		} else {
			if i != 0 {
				if a[i-1] == ' ' {
					continue
				} else {
					b = append(b, a[i])
				}
			} else {
				continue
			}
		}
	}
	return b
}

func reverse(a []byte, l, r int) {
	leftP := l
	rightP := r
	for leftP <= rightP {
		t := a[leftP]
		a[leftP] = a[rightP]
		a[rightP] = t
		leftP++
		rightP--
	}
}
