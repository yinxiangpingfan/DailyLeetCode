package hashmap

func twoSum(nums []int, target int) []int {
	a := make(map[int]int)
	for k, v := range nums {
		a[target-v] = k
	}
	for k, v := range nums {
		if v, ok := a[v]; ok {
			if k != v {
				return []int{
					v,
					k,
				}
			}
		}
	}
	return []int{}
}
