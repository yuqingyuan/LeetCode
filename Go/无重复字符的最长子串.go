package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var i, max int
	m := make(map[byte]int)
	for {
		if i >= len(s) {
			break
		}
		if index, ok := m[s[i]]; ok {
			num := len(m)
			if num > max {
				max = num
			}
			m = make(map[byte]int)
			i = 0
			s = s[index+1:]
		} else {
			m[s[i]] = i
			i++
		}
	}
	if len(m) > max {
		max = len(m)
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

