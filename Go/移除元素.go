package main

import "fmt"

func removeElement(nums []int, val int) int {
	var index int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	fmt.Println(nums)
	return index
}

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

