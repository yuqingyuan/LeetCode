package main

import (
	"fmt"
)

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return false
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else {
			prior := stack[len(stack)-1]
			if s[i] == ')' {
				if prior != '(' {
					return false
				}
			} else if s[i] == ']' {
				if prior != '[' {
					return false
				}
			} else if s[i] == '}' {
				if prior != '{' {
					return false
				}
			} else {
				stack = append(stack, s[i])
				continue
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("(("))
}

