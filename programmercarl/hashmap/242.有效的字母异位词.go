package hashmap

func isAnagram(s string, t string) bool {
	lens := len(s)
	a := make([]int, 26)
	if len(t) != lens {
		return false
	}
	s1 := []byte(s)
	t1 := []byte(t)
	for _, i := range s1 {
		a[i-'a'] += 1
	}
	for _, i := range t1 {
		a[i-'a']--
		if a[i-'a'] < 0 {
			return false
		}
	}
	for _, v := range a {
		if v != 0 {
			return false
		}
	}
	return true
}
