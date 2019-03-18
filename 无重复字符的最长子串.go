package main

func lengthOfLongestSubstring(s string) int {
	var max int
	start := 0
	end := start + 1
	for {
		if end > len(s) {
			break
		}
		flag := isUnique(s[start:end])
		if flag {
			if end-start > max {
				max = end - start
			}
			end++
		} else {
			if end-start-1 > max {
				max = end - start - 1
			}
			start += 1
			end = start + 1
		}
	}
	return max
}

func isUnique(s string) bool {
	flag := true
	m := make(map[rune]int)
	for index, v := range s {
		if len(m) == 0 {
			m[v] = index
		} else {
			if _, ok := m[v]; ok {
				flag = false
				break
			} else {
				m[v] = index
				flag = true
			}
		}
	}
	return flag
}

func main() {
	lengthOfLongestSubstring("alqebriavxoo")
}

