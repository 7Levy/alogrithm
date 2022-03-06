package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	n := len(nums)
	for i := 0; i < n-3; i++ {
		n1 := nums[i]
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			n2 := nums[j]
			if j > i+1 && n2 == nums[j-1] {
				continue
			}
			l := j + 1
			r := n - 1

			for l < r {
				n3 := nums[l]
				n4 := nums[r]
				sum := n1 + n2 + n3 + n4
				if sum < target {
					l++
				} else if sum > target {
					r--
				} else {
					res = append(res, []int{n1, n2, n3, n4})
					for l < r && n3 == nums[l] {
						l++
					}
					for l < r && n4 == nums[r] {
						r--
					}
				}
			}
		}
	}
	return res
}
