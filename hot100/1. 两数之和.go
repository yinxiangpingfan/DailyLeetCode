package hot100

func twoSum(nums []int, target int) []int {
	rs := make([]int, 0)
	a := make(map[int]int)
	for k, v := range nums {
		if kv, ok := a[v]; ok {
			if v+v == target {
				return []int{
					kv,
					k,
				}
			}
		} else {
			if kv1, ok1 := a[target-v]; ok1 {
				return []int{
					k,
					kv1,
				}
			} else {
				a[v] = k
			}
		}
	}
	return rs
}
