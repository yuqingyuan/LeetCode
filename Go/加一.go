package main

import "fmt"

func plusOne(digits []int) []int {
	length := len(digits)
	var flag int
	for i := length - 1; i >= 0; i-- {
		if flag = (digits[i] + 1) / 10; flag == 1 {
			digits[i] = 0
		} else {
			digits[i] += 1
			break
		}
	}
	if flag == 1 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 0}))
}

