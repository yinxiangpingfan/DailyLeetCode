package hashmap

func intersection(nums1 []int, nums2 []int) []int {
	a := make(map[int]interface{})
	b := make(map[int]interface{})
	returnSlice := make([]int, 0)
	for _, v := range nums1 {
		a[v] = struct{}{}
	}
	for _, v := range nums2 {
		b[v] = struct{}{}
	}
	if len(a) > len(b) {
		for k, _ := range b {
			if _, ok := a[k]; ok {
				returnSlice = append(returnSlice, k)
			}
		}
	} else {
		for k, _ := range a {
			if _, ok := b[k]; ok {
				returnSlice = append(returnSlice, k)
			}
		}
	}
	return returnSlice
}
