package list

func isHappy(n int) bool {
	a := make(map[int]interface{})
	a[n] = struct{}{}
	for {
		count := 0
		for n != 0 {
			ge := n % 10
			count += ge * ge
			n /= 10
		}
		n = count
		if n == 1 {
			return true
		}
		if _, ok := a[n]; ok {
			return false
		} else {
			a[n] = struct{}{}
		}
	}
}
