package hashmap

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	returnSlice := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for k, v := range nums {
		if v > 0 {
			break
		}
		if k >= 1 && v == nums[k-1] {
			continue
		}
		target := 0 - v
		leftP := k + 1
		rightP := len(nums) - 1
		for leftP < rightP {
			if leftP > k+1 && nums[leftP-1] == nums[leftP] {
				leftP++
				continue
			}
			if rightP < len(nums)-1 && nums[rightP+1] == nums[rightP] {
				rightP--
				continue
			}
			if leftP >= rightP {
				break
			}
			if nums[leftP]+nums[rightP] < target {
				leftP++
			} else if nums[leftP]+nums[rightP] == target {
				returnSlice = append(returnSlice, []int{
					v,
					nums[leftP],
					nums[rightP],
				})
				leftP++
				rightP--
			} else {
				rightP--
			}
		}
	}
	return returnSlice
}
