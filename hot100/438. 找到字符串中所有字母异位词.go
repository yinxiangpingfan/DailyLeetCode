package hot100

func findAnagrams(s string, p string) []int {
	leftP := 0
	right := leftP + len(p) - 1
	b := []byte(s)
	pp := []byte(p)
	var f [26]int
	for i := 0; i < len(p); i++ {
		f[pp[i]-'a']++
	}
	ri := make([]int, 0)
	var a [26]int
	for right < len(s) {
		if leftP == 0 {
			for i := leftP; i <= right; i++ {
				a[b[i]-'a']++
			}
		} else {
			a[b[right]-'a']++
		}
		if f == a {
			ri = append(ri, leftP)
		}
		a[b[leftP]-'a']--
		leftP++
		right++
	}
	return ri
}
