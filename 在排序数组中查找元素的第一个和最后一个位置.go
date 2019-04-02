package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	index := binarySearch(nums, 0, len(nums)-1, target)
	if index == -1 {
		return []int{-1, -1}
	} else {
		start := index
		end := index
		for i := index - 1; i >= 0; i-- {
			if nums[i] == nums[index] {
				start = i
			}
		}
		for j := index + 1; j < len(nums); j++ {
			if nums[j] == nums[index] {
				end = j
			}
		}
		return []int{start, end}
	}
}

func binarySearch(nums []int, start int, end int, target int) int {
	if start > end {
		return -1
	}
	mid := start + (end-start)/2
	if nums[mid] > target {
		return binarySearch(nums, start, mid-1, target)
	} else if nums[mid] < target {
		return binarySearch(nums, mid+1, end, target)
	} else {
		return mid
	}
}

func main() {
	nums := []int{1, 2, 3, 3, 3, 3, 4, 5, 9}
	fmt.Println(searchRange(nums, 3))
}

