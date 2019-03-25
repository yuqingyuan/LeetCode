package main

import "fmt"

func maxArea(height []int) int {
	var max, i, j int
	j = len(height) - 1
	for {
		if i == j {
			break
		}
		var v int
		if height[i] > height[j] {
			v = height[j] * (j - i)
			j--
		} else {
			v = height[i] * (j - i)
			i++
		}
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(maxArea([]int{1, 8}))
}

