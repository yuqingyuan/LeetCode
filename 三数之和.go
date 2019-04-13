package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	if len(nums) <= 2 {
		return res
	}

	for i := 0; i <= len(nums)-3; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			remain := -nums[i]
			left := i + 1
			right := len(nums) - 1
			for left < right {
				if nums[left]+nums[right] == remain {
					var temp []int
					temp = append(temp, nums[i], nums[left], nums[right])
					res = append(res, temp)
					for {
						left += 1
						if left >= right || nums[left] != nums[left-1] {
							break
						}
					}
					for {
						right -= 1
						if left >= right || nums[right] != nums[right+1] {
							break
						}
					}
				} else if nums[left]+nums[right] < remain {
					for {
						left += 1
						if left >= right || nums[left] != nums[left-1] {
							break
						}
					}
				} else {
					for {
						right -= 1
						if left >= right || nums[right] != nums[right+1] {
							break
						}
					}
				}
			}
		}
	}

	return res
}

func main() {
	fmt.Print(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

