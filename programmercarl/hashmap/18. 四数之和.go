package hashmap

import "sort"

func fourSum(nums []int, target int) [][]int {
	returnSlice := make([][]int, 0)
	if len(nums) < 4 {
		return nil
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i, v1 := range nums {
		if i > 0 && v1 == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			v2 := nums[j]
			if j > i+1 && v2 == nums[j-1] {
				continue
			}
			aim := target - v1 - v2
			leftP := j + 1
			rightP := len(nums) - 1
			for leftP < rightP {
				if leftP > j+1 && nums[leftP] == nums[leftP-1] {
					leftP++
					continue
				}
				if rightP < len(nums)-1 && nums[rightP] == nums[rightP+1] {
					rightP--
					continue
				}
				if nums[leftP]+nums[rightP] < aim {
					leftP++
				} else if nums[leftP]+nums[rightP] > aim {
					rightP--
				} else {
					returnSlice = append(returnSlice, []int{
						nums[leftP],
						nums[rightP],
						v1, v2,
					})
					leftP++
					rightP--
				}
			}
		}
	}
	return returnSlice
}
