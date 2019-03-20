package main

import (
	"fmt"
	"strings"
)

func longestPalindrome(s string) string {
	res := []rune{'#'}
	for _, word := range s {
		res = append(res, word)
		res = append(res, '#')
	}
	var nums int
	var raw []rune
	for i := 0; i < len(res); i++ {
		len := expand(res, i)
		if len > nums {
			nums = len
			raw = res[i-(nums-1)/2 : i+(nums-1)/2+1]
		}
	}
	result := []string{}
	for _, value := range raw {
		if value != '#' {
			result = append(result, string(value))
		}
	}
	return strings.Join(result, "")
}

func expand(s []rune, index int) int {
	left := index
	right := index
	for {
		if left >= 0 && right < len(s) && s[left] == s[right] {
			left = left - 1
			right = right + 1
		} else {
			break
		}
	}
	return right - left - 1
}

func main() {
	fmt.Println(longestPalindrome("aabb"))
}

