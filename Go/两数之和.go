package main

import "fmt"

// func twoSum(nums []int, target int) []int {
// 	for index, v := range nums {
// 		for i := index + 1; i < len(nums); i++ {
// 			if result := v + nums[i]; result == target {
// 				return []int{index, i}
// 			}
// 		}
// 	}
// 	return []int{}
// }

//最佳解法
func twoSum(nums []int, target int) []int {
	tables := make(map[int]int)
	for index, value := range nums {
		tmp := target - value
		if i, ok := tables[tmp]; ok {
			return []int{i, index}
		}
		tables[value] = index
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

