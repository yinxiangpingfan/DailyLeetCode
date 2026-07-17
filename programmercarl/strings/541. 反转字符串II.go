package strings

func reverseStr(s string, k int) string {
	a := []byte(s)
	lens := len(s)
	for i := 0; i < lens/(2*k); i++ {
		leftP := i * 2 * k
		rightP := leftP + k - 1
		for leftP <= rightP {
			t := a[leftP]
			a[leftP] = a[rightP]
			a[rightP] = t
			leftP++
			rightP--
		}
	}
	leave := lens % (2 * k)
	if leave < k {
		leftP := lens - leave
		rightP := lens - 1
		for leftP <= rightP {
			t := a[leftP]
			a[leftP] = a[rightP]
			a[rightP] = t
			leftP++
			rightP--
		}
	} else {
		leftP := lens - leave
		rightP := leftP + k - 1
		for leftP <= rightP {
			t := a[leftP]
			a[leftP] = a[rightP]
			a[rightP] = t
			leftP++
			rightP--
		}
	}
	return string(a)
}
