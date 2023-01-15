package twoSum

import (
	"sort"
)

func twoSum(nums []int, target int) []int {
	smallerNums := make([]int, len(nums))
	copy(smallerNums, nums)
	// O(n log n)
	sort.Ints(smallerNums)
	values := []int{0, 0}
	i := 0
	j := len(smallerNums) - 1
	for i < j {
		sum := smallerNums[i] + smallerNums[j]
		if sum < target {
			i++
		} else if sum > target {
			j--
		} else {
			values[0] = smallerNums[i]
			values[1] = smallerNums[j]
			break
		}
	}

	indices := []int{-1, -1}
	for idx, v := range nums {
		if v == values[0] && indices[0] == -1 {
			indices[0] = idx
		} else if v == values[1] && indices[1] == -1 {
			indices[1] = idx
		}
		if indices[0] != -1 && indices[1] != -1 {
			break
		}
	}
	sort.Ints(indices)
	return indices
}
