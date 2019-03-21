package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	var s string
	var isNeg bool = false
	if x < 0 {
		isNeg = true
	}
	s = strconv.Itoa(x)
	if isNeg {
		s = s[1:]
	}
	var flag bool = false
	result := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 0 && !flag {
			flag = true
		}
		if flag {
			result = append(result, string(s[i]))
		}
	}
	if isNeg {
		result = append([]string{"-"}, result...)
	}
	v, _ := strconv.Atoi(strings.Join(result, ""))
	if v > math.MaxInt32 || v < math.MinInt32 {
		return 0
	}
	return v
}

func main() {
	fmt.Println(reverse(1563847412))
}

