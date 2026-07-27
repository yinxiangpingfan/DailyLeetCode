package doubleptr

import "sort"

func fourSum(nums []int, target int) [][]int {
	r := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			leftP := j + 1
			rightP := len(nums) - 1
			for leftP < rightP {
				if leftP != j+1 && nums[leftP] == nums[leftP-1] {
					leftP++
					continue
				}
				if rightP != len(nums)-1 && nums[rightP] == nums[rightP+1] {
					rightP--
					continue
				}
				sum := nums[leftP] + nums[rightP] + nums[i] + nums[j]
				if sum == target {
					t := []int{
						nums[leftP],
						nums[rightP],
						nums[i],
						nums[j],
					}
					leftP++
					rightP--
					r = append(r, t)
				} else if sum < target {
					leftP++
				} else {
					rightP--
				}
			}
		}
	}
	return r
}
