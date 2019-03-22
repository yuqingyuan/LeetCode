package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	length := len(s)
	if length == 1 {
		return true
	}
	s1 := s[:length/2]
	var s2 string
	if length%2 == 0 {
		s2 = s[length/2:]
	} else {
		s2 = s[length/2+1:]
	}
	s2 = Reverse(s2)

	if strings.Compare(s1, s2) == 0 {
		return true
	} else {
		return false
	}
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	fmt.Println(isPalindrome(12345654321))
}

