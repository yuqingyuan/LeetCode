package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) < numRows {
		return s
	}
	result := []string{}
	gap := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		result = append(result, string(s[i]))
		start := i
		for {
			start += gap
			if i != 0 && i != numRows-1 {
				if start < len(s) {
					result = append(result, string(s[start-i*2]))
					result = append(result, string(s[start]))
				} else {
					if start-i*2 < len(s) {
						result = append(result, string(s[start-i*2]))
					} else {
						break
					}
				}
			} else {
				if start >= len(s) {
					break
				}
				result = append(result, string(s[start]))
			}
		}
	}
	return strings.Join(result, "")
}

func main() {
	fmt.Println(convert("A", 2))
}

