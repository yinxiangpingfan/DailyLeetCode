package hot100

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	rs := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			if l != i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			if r != len(nums)-1 && nums[r] == nums[r+1] {
				r--
				continue
			}
			if nums[l]+nums[r]+nums[i] == 0 {
				t := []int{
					nums[i],
					nums[l],
					nums[r],
				}
				rs = append(rs, t)
				l++
				r--
			} else if nums[l]+nums[r]+nums[i] < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return rs
}
