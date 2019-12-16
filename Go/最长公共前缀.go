package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	s1 := strs[0]
	for i := 1; i < len(strs); i++ {
		for {
			if len(s1) == 0 {
				return ""
			}
			if strings.HasPrefix(strs[i], s1) {
				break
			} else {
				s1 = s1[:len(s1)-1]
			}
		}
	}

	return s1
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"abca", "abc"}))
}

