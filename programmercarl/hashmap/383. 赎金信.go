package hashmap

func canConstruct(ransomNote string, magazine string) bool {
	a := make(map[byte]int)
	mb := []byte(magazine)
	for _, v := range mb {
		a[v] += 1
	}
	mb = []byte(ransomNote)
	for _, v := range mb {
		a[v] -= 1
		if a[v] < 0 {
			return false
		}
	}
	return true
}
