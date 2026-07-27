package doubleptr

import "sort"

func threeSum(nums []int) [][]int {
	returnSlice := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		leftP := i + 1
		rightP := len(nums) - 1
		for leftP < rightP {
			if leftP != i+1 && nums[leftP] == nums[leftP-1] {
				leftP++
				continue
			}
			if rightP != len(nums)-1 && nums[rightP] == nums[rightP+1] {
				rightP--
				continue
			}
			if leftP >= rightP {
				break
			}
			if nums[i]+nums[leftP]+nums[rightP] < 0 {
				leftP++
			} else if nums[i]+nums[leftP]+nums[rightP] == 0 {
				t := []int{
					nums[leftP],
					nums[rightP],
					nums[i],
				}
				returnSlice = append(returnSlice, t)
				leftP++
				rightP--
			} else {
				rightP--
			}
		}
	}
	return returnSlice
}
