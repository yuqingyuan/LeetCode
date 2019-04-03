package main

import "fmt"

func searchInsert(nums []int, target int) int {
	index := binarySearch(nums, 0, len(nums)-1, target)
	return index
}

func binarySearch(nums []int, start int, end int, target int) int {
	if start > end {
		return start
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
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 7))
}

