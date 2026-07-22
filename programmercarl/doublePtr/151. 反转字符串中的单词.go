package doubleptr

func reverseWords(s string) string {
	s1 := []byte(s)
	remove(s1, 0, len(s1)-1)

	s1 = deleteN(s1)

	leftP := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == ' ' {
			remove(s1, leftP, i-1)
			leftP = i + 1
		}
	}
	remove(s1, leftP, len(s1)-1)
	return string(s1)
}

func deleteN(s []byte) []byte {
	returnByte := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			returnByte = append(returnByte, s[i])
		} else {
			if i+1 <= len(s)-1 && s[i+1] != ' ' {
				returnByte = append(returnByte, s[i])
			} else {
				continue
			}
		}
	}
	if returnByte[0] == ' ' {
		return returnByte[1:]
	}
	return returnByte
}

func remove(s []byte, leftP, rightP int) {
	for leftP <= rightP {
		t := s[leftP]
		s[leftP] = s[rightP]
		s[rightP] = t
		leftP++
		rightP--
	}
}
