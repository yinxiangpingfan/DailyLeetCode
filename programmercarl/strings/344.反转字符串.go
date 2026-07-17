package strings

func reverseString(s []byte) {
	leftP := 0
	rightP := len(s) - 1
	for leftP <= rightP {
		t := s[leftP]
		s[leftP] = s[rightP]
		s[rightP] = t
		leftP++
		rightP--
	}
}
